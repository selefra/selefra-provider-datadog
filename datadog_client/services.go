package datadog_client

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/selefra/selefra-provider-datadog/constants"

	datadogV1 "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	datadogV2 "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

var (
	defaultHTTPRetryDuration = 5 * time.Second
	defaultHTTPRetryTimeout  = 60 * time.Second
	rateLimitResetHeader     = constants.XRatelimitReset
)

type CustomTransport struct {
	defaultTransport  http.RoundTripper
	httpRetryDuration time.Duration
	httpRetryTimeout  time.Duration
}

type CustomTransportOptions struct {
	Timeout *time.Duration
}

func (t *CustomTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	var ccancel context.CancelFunc
	ctx := req.Context()
	if _, set := ctx.Deadline(); !set {
		ctx, ccancel = context.WithTimeout(ctx, t.httpRetryTimeout)
		defer ccancel()
	}

	retryCount := 0
	for {
		newRequest := t.copyRequest(req)
		resp, respErr := t.defaultTransport.RoundTrip(newRequest)

		if resp != nil {
			localVarBody, _ := io.ReadAll(resp.Body)
			resp.Body.Close()
			resp.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
		}
		if respErr != nil {
			return resp, respErr
		}

		retryDuration, retry := t.retryRequest(resp)
		if !retry {
			return resp, respErr
		}

		if retryDuration == nil {
			newRetryDurationVal := time.Duration(retryCount) * t.httpRetryDuration
			retryDuration = &newRetryDurationVal
		}

		select {
		case <-ctx.Done():
			return resp, respErr
		case <-time.After(*retryDuration):
			retryCount++
			continue
		}
	}
}

func (t *CustomTransport) copyRequest(r *http.Request) *http.Request {
	newRequest := *r

	if r.Body == nil || r.Body == http.NoBody {
		return &newRequest
	}

	body, _ := r.GetBody()
	newRequest.Body = body

	return &newRequest
}

func (t *CustomTransport) retryRequest(response *http.Response) (*time.Duration, bool) {
	if v := response.Header.Get(rateLimitResetHeader); v != constants.Constants_0 && response.StatusCode == 429 {
		vInt, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return nil, true
		}
		retryDuration := time.Duration(vInt) * time.Second
		return &retryDuration, true
	}

	if response.StatusCode >= 500 {
		return nil, true
	}

	return nil, false
}

func NewCustomTransport(t http.RoundTripper, opt CustomTransportOptions) *CustomTransport {

	if t == nil {
		t = http.DefaultTransport
	}

	ct := CustomTransport{
		defaultTransport:  t,
		httpRetryDuration: defaultHTTPRetryDuration,
	}

	if opt.Timeout != nil {
		ct.httpRetryTimeout = *opt.Timeout
	} else {
		ct.httpRetryTimeout = defaultHTTPRetryTimeout
	}

	return &ct
}

func (t *CustomTransport) DefaultBackoff(min, max time.Duration, attemptNum int, resp *http.Response) time.Duration {
	if resp != nil {
		if resp.StatusCode == http.StatusTooManyRequests || resp.StatusCode == http.StatusServiceUnavailable {
			if s, ok := resp.Header[constants.RetryAfter]; ok {
				if sleep, err := strconv.ParseInt(s[0], 10, 64); err == nil {
					return time.Second * time.Duration(sleep)
				}
			}
		}
	}

	mult := math.Pow(2, float64(attemptNum)) * float64(min)
	sleep := time.Duration(mult)
	if float64(sleep) != mult || sleep > max {
		sleep = max
	}
	return sleep
}

func coverConfig(ctx context.Context, config *Config) (context.Context, error) {
	apiKey := os.Getenv(constants.DDCLIENTAPIKEY)
	appKey := os.Getenv(constants.DDCLIENTAPPKEY)
	apiURL := constants.Httpsapidatadoghqcom

	if config.ApiKey != constants.Constants_1 {
		apiKey = config.ApiKey
	}
	if config.AppKey != constants.Constants_2 {
		appKey = config.AppKey
	}
	if config.ApiUrl != constants.Constants_3 {
		apiURL = config.ApiUrl
	}
	if apiKey == constants.Constants_4 {
		return ctx, errors.New(constants.Apikeymustbeconfigured)
	}
	if appKey == constants.Constants_5 {
		return ctx, errors.New(constants.Appkeymustbeconfigured)
	}

	ctx = context.WithValue(ctx, datadogV1.ContextAPIKeys,
		map[string]datadogV1.APIKey{
			constants.ApiKeyAuth: {Key: apiKey},
			constants.AppKeyAuth: {Key: appKey},
		},
	)

	if apiURL != constants.Constants_6 {
		parsedAPIURL, parseErr := url.Parse(apiURL)
		if parseErr != nil {
			return ctx, fmt.Errorf(`invalid API URL : %v`, parseErr)
		}
		if parsedAPIURL.Host == constants.Constants_7 || parsedAPIURL.Scheme == constants.Constants_8 {
			return ctx, fmt.Errorf(`missing protocol or host : %v`, apiURL)
		}

		strings.Split(parsedAPIURL.Host, constants.Constants_9)

		ctx = context.WithValue(ctx, datadogV1.ContextServerIndex, 1)
		ctx = context.WithValue(ctx,
			datadogV1.ContextServerVariables,
			map[string]string{
				constants.Name:     parsedAPIURL.Host,
				constants.Protocol: parsedAPIURL.Scheme,
			})
	}
	return ctx, nil
}

func V1(ctx context.Context, config *Config) (context.Context, *datadogV1.APIClient, error) {
	ctx, err := coverConfig(ctx, config)
	if err != nil {
		return ctx, nil, err
	}

	httpClientV1 := http.DefaultClient
	ctOptions := CustomTransportOptions{}
	timeout := time.Duration(int64(60)) * time.Second
	ctOptions.Timeout = &timeout
	httpClientV1.Transport = NewCustomTransport(httpClientV1.Transport, ctOptions)

	configuration := datadogV1.NewConfiguration()
	configuration.HTTPClient = httpClientV1
	apiClient := datadogV1.NewAPIClient(configuration)
	return ctx, apiClient, nil
}

func V2(ctx context.Context, config *Config) (context.Context, *datadogV2.APIClient, *datadogV2.Configuration, error) {
	ctx, err := coverConfig(ctx, config)
	if err != nil {
		return ctx, nil, nil, err
	}
	ctx = context.WithValue(
		ctx,
		datadogV2.ContextServerVariables,
		map[string]string{constants.BasePath: constants.V},
	)
	httpClientV2 := http.DefaultClient
	ctOptions := CustomTransportOptions{}
	timeout := time.Duration(int64(60)) * time.Second
	ctOptions.Timeout = &timeout
	httpClientV2.Transport = NewCustomTransport(httpClientV2.Transport, ctOptions)

	configuration := datadogV2.NewConfiguration()
	configuration.HTTPClient = httpClientV2
	apiClient := datadogV2.NewAPIClient(configuration)
	return ctx, apiClient, configuration, nil
}
