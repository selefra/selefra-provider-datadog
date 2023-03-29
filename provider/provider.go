package provider

import (
	"context"
	"os"

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
				var datadogConfig datadog_client.Config
				err := config.Unmarshal(&datadogConfig)
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorMsg(constants.Analysisconfigerrs, err.Error())
				}

				if datadogConfig.ApiKey == "" {
					datadogConfig.ApiKey = os.Getenv("DATADOG_API_KEY")
				}

				if datadogConfig.ApiKey == "" {
					return nil, schema.NewDiagnostics().AddErrorMsg("missing ApiKey in configuration")
				}

				if datadogConfig.AppKey == "" {
					datadogConfig.AppKey = os.Getenv("DATADOG_APP_KEY")
				}

				if datadogConfig.AppKey == "" {
					return nil, schema.NewDiagnostics().AddErrorMsg("missing AppKey in configuration")
				}

				if datadogConfig.ApiUrl == "" {
					datadogConfig.ApiUrl = os.Getenv("DATADOG_API_URL")
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
				return `# api_key: <Your Datadog Api Key>
# app_key: <Your Datadog App Key>
# api_url: <Your Datadog Api Url>`
			},
			Validation: func(ctx context.Context, config *viper.Viper) *schema.Diagnostics {
				var datadogConfig datadog_client.Config
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
