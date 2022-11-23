package tables

import (
	"context"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"github.com/selefra/selefra-provider-datadog/datadog_client"
	"github.com/selefra/selefra-provider-datadog/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableDatadogDashboardGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableDatadogDashboardGenerator{}

func (x *TableDatadogDashboardGenerator) GetTableName() string {
	return "datadog_dashboard"
}

func (x *TableDatadogDashboardGenerator) GetTableDescription() string {
	return ""
}

func (x *TableDatadogDashboardGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableDatadogDashboardGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableDatadogDashboardGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			ctx, apiClient, err := datadog_client.V1(ctx, taskClient.(*datadog_client.Client).Config)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			resp, _, err := apiClient.DashboardsApi.ListDashboards(ctx, datadog.ListDashboardsOptionalParameters{})
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			for _, dashboard := range resp.GetDashboards() {
				resultChannel <- dashboard

			}

			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

func (x *TableDatadogDashboardGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableDatadogDashboardGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("url").ColumnType(schema.ColumnTypeString).Description("URL of the dashboard.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("widgets").ColumnType(schema.ColumnTypeJSON).Description("List of widgets to display on the dashboard.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("layout_type").ColumnType(schema.ColumnTypeString).Description("Layout type of the dashboard. Can be on of \"free\" or \"ordered\".").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("reflow_type").ColumnType(schema.ColumnTypeString).Description("Reflow type for a new dashboard layout dashboard. If set to 'fixed', the dashboard expects all widgets to have a layout, and if it's set to 'auto', widgets should not have layouts.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the dashboard.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("template_variable_presets").ColumnType(schema.ColumnTypeJSON).Description("List of template variables saved views.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeTimestamp).Description("Creation date of the dashboard.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("author_handle").ColumnType(schema.ColumnTypeString).Description("Identifier of the dashboard author.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Description("Description of the dashboard.").
			Extractor(column_value_extractor.StructSelector("Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_read_only").ColumnType(schema.ColumnTypeBool).Description("Indicates if the dashboard is read-only. If True, only the author and admins can make changes to it.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("modified_at").ColumnType(schema.ColumnTypeTimestamp).Description("Modification time of the dashboard.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("restricted_roles").ColumnType(schema.ColumnTypeJSON).Description("A list of role identifiers. Only the author and users associated with at least one of these roles can edit this dashboard. Overrides the `is_read_only` property if both are present.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("template_variables").ColumnType(schema.ColumnTypeJSON).Description("List of template variables for this dashboard.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("Dashboard identifier.").Build(),
	}
}

func (x *TableDatadogDashboardGenerator) GetSubTables() []*schema.Table {
	return nil
}
