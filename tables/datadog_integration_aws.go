package tables

import (
	"context"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"github.com/selefra/selefra-provider-datadog/datadog_client"
	"github.com/selefra/selefra-provider-datadog/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
)

type TableDatadogIntegrationAwsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableDatadogIntegrationAwsGenerator{}

func (x *TableDatadogIntegrationAwsGenerator) GetTableName() string {
	return "datadog_integration_aws"
}

func (x *TableDatadogIntegrationAwsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableDatadogIntegrationAwsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableDatadogIntegrationAwsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableDatadogIntegrationAwsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			ctx, apiClient, err := datadog_client.V1(ctx, taskClient.(*datadog_client.Client).Config)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			opts := datadog.ListAWSAccountsOptionalParameters{}

			resp, _, err := apiClient.AWSIntegrationApi.ListAWSAccounts(ctx, opts)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			for _, account := range resp.GetAccounts() {
				resultChannel <- account

			}

			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

func (x *TableDatadogIntegrationAwsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableDatadogIntegrationAwsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("account_specific_namespace_rules").ColumnType(schema.ColumnTypeJSON).Description("An object, that enables or disables metric collection for specific AWS namespaces for this AWS account only.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("excluded_regions").ColumnType(schema.ColumnTypeJSON).Description("An array of AWS regions to exclude from metrics collection.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("filter_tags").ColumnType(schema.ColumnTypeJSON).Description("List of tags (in the form 'key:value') that define a filter which is used when collecting EC2 or Lambda resources. These key:value pairs can be used to both whitelist and blacklist tags.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("host_tags").ColumnType(schema.ColumnTypeJSON).Description("Array of tags (in the form `key:value`) to add to all hosts and metrics reporting through this integration.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).Description("Your AWS Account ID without dashes.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cspm_resource_collection_enabled").ColumnType(schema.ColumnTypeBool).Description("Whether Datadog collects cloud security posture management resources from your AWS account. This includes additional resources not covered under the general `resource_collection`.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("secret_access_key").ColumnType(schema.ColumnTypeString).Description("Your AWS secret access key. Only required if your AWS account is a GovCloud or China account.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("metrics_collection_enabled").ColumnType(schema.ColumnTypeBool).Description("Whether Datadog collects metrics for this AWS account.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resource_collection_enabled").ColumnType(schema.ColumnTypeBool).Description("Whether Datadog collects a standard set of resources from your AWS account.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("role_name").ColumnType(schema.ColumnTypeString).Description("Your Datadog role delegation name.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("access_key_id").ColumnType(schema.ColumnTypeString).Description("Your AWS access key ID. Only required if your AWS account is a GovCloud or China account.").Build(),
	}
}

func (x *TableDatadogIntegrationAwsGenerator) GetSubTables() []*schema.Table {
	return nil
}
