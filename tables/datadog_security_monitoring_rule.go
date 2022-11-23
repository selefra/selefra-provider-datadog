package tables

import (
	"context"

	"github.com/DataDog/datadog-api-client-go/api/v2/datadog"
	"github.com/selefra/selefra-provider-datadog/datadog_client"
	"github.com/selefra/selefra-provider-datadog/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
)

type TableDatadogSecurityMonitoringRuleGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableDatadogSecurityMonitoringRuleGenerator{}

func (x *TableDatadogSecurityMonitoringRuleGenerator) GetTableName() string {
	return "datadog_security_monitoring_rule"
}

func (x *TableDatadogSecurityMonitoringRuleGenerator) GetTableDescription() string {
	return ""
}

func (x *TableDatadogSecurityMonitoringRuleGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableDatadogSecurityMonitoringRuleGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableDatadogSecurityMonitoringRuleGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			ctx, apiClient, _, err := datadog_client.V2(ctx, taskClient.(*datadog_client.Client).Config)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			opts := datadog.ListSecurityMonitoringRulesOptionalParameters{
				PageSize:   datadog.PtrInt64(100),
				PageNumber: datadog.PtrInt64(0),
			}

			count := int64(0)
			for {
				resp, _, err := apiClient.SecurityMonitoringApi.ListSecurityMonitoringRules(ctx, opts)
				if err != nil {

					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}

				for _, securityMonitoringRule := range resp.GetData() {
					count++
					resultChannel <- securityMonitoringRule

				}

				if count >= resp.Meta.Page.GetTotalCount() {
					return schema.NewDiagnosticsErrorPullTable(task.Table, nil)
				}
				opts.WithPageNumber(*opts.PageNumber + 1)
			}

		},
	}
}

func (x *TableDatadogSecurityMonitoringRuleGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableDatadogSecurityMonitoringRuleGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("message").ColumnType(schema.ColumnTypeString).Description("Message for generated signals.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("options").ColumnType(schema.ColumnTypeJSON).Description("Additional options for security monitoring rules.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("The ID of the rule.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("update_author_id").ColumnType(schema.ColumnTypeString).Description("User ID of the user who updated the rule.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeString).Description("When the rule was created, timestamp in milliseconds.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_default").ColumnType(schema.ColumnTypeBool).Description("Whether the rule is included by default.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("version").ColumnType(schema.ColumnTypeInt).Description("The version of the rule.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Description("Tags for generated signals.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_enabled").ColumnType(schema.ColumnTypeBool).Description("Whether the rule is enabled.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("queries").ColumnType(schema.ColumnTypeString).Description("Queries for selecting logs which are part of the rule.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).Description("The security monitoring rule type.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cases").ColumnType(schema.ColumnTypeJSON).Description("Cases for generating signals.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("The name of the rule.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("creation_author_id").ColumnType(schema.ColumnTypeString).Description("User ID of the user who created the rule.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("has_extended_title").ColumnType(schema.ColumnTypeBool).Description("Whether the notifications include the triggering group-by values in their title.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_deleted").ColumnType(schema.ColumnTypeBool).Description("Whether the rule has been deleted.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("filters").ColumnType(schema.ColumnTypeJSON).Description("Additional queries to filter matched events before they are processed.").Build(),
	}
}

func (x *TableDatadogSecurityMonitoringRuleGenerator) GetSubTables() []*schema.Table {
	return nil
}
