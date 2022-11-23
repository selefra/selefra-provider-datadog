package provider

import (
	"context"

	"github.com/selefra/selefra-provider-datadog/constants"

	"github.com/selefra/selefra-provider-datadog/datadog_client"

	"github.com/selefra/selefra-provider-sdk/provider"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/spf13/viper"
)

var Version = "v0.0.1"

func GetProvider() *provider.Provider {
	return &provider.Provider{
		Name:      constants.Datadog,
		Version:   Version,
		TableList: GenTables(),
		ClientMeta: schema.ClientMeta{
			InitClient: func(ctx context.Context, clientMeta *schema.ClientMeta, config *viper.Viper) ([]any, *schema.Diagnostics) {
				var datadogConfig datadog_client.Configs
				err := config.Unmarshal(&datadogConfig)
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorMsg(constants.Analysisconfigerrs, err.Error())
				}
				if len(datadogConfig.Providers) == 0 {
					datadogConfig.Providers = append(datadogConfig.Providers, datadog_client.Config{})
				}

				clients, err := datadog_client.NewClients(datadogConfig)

				if err != nil {
					clientMeta.ErrorF(constants.Newclientserrs, err.Error())
					return nil, schema.NewDiagnostics().AddError(err)
				}

				if len(clients) == 0 {
					return nil, schema.NewDiagnostics().AddErrorMsg(constants.Accountinformationnotfound)
				}

				res := make([]interface{}, 0, len(clients))
				for i := range clients {
					res = append(res, clients[i])
				}
				return res, nil
			},
		},
		ConfigMeta: provider.ConfigMeta{
			GetDefaultConfigTemplate: func(ctx context.Context) string {
				return `##  Optional, Repeated. Add an accounts block for every account you want to assume-role into and fetch data from.
#accounts:
#  - api_key: # (required) - API keys are unique to an organization. An API key is required by the Datadog Agent to submit metrics and events to Datadog. Get an API key. May alternatively be set via the DD_CLIENT_API_KEY environment variable.
#    app_key: # (required) - Application keys in conjunction with organization’s API key, give users access to Datadog’s programmatic API. Application keys are associated with the user account that created them and have the permissions and capabilities of the user who created them. Get an application key. May alternatively be set via the DD_CLIENT_APP_KEY environment variable.
#    api_url: # (optional) - The API URL used for all requests. Defaults to "https://api.datadoghq.com/". If working with the EU version, this should be changed to "https://api.datadoghq.eu/".
`
			},
			Validation: func(ctx context.Context, config *viper.Viper) *schema.Diagnostics {
				var datadogConfig datadog_client.Configs
				err := config.Unmarshal(&datadogConfig)
				if err != nil {
					return schema.NewDiagnostics().AddErrorMsg(constants.Analysisconfigerrs, err.Error())
				}
				return nil
			},
		},
		TransformerMeta: schema.TransformerMeta{
			DefaultColumnValueConvertorBlackList: []string{
				constants.Constants_10,
				constants.NA,
				constants.Notsupported,
			},
			DataSourcePullResultAutoExpand: true,
		},
		ErrorsHandlerMeta: schema.ErrorsHandlerMeta{

			IgnoredErrors: []schema.IgnoredError{schema.IgnoredErrorOnSaveResult},
		},
	}
}
