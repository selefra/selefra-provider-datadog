package tables

import (
	"context"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"

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

			ctx, apiClient, _, err := datadog_client.Server(ctx, taskClient.(*datadog_client.Client).Config)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			api := datadogV1.NewDashboardsApi(apiClient)

			resp, _, err := api.ListDashboards(ctx)

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
		table_schema_generator.NewColumnBuilder().ColumnName("url").ColumnType(schema.ColumnTypeString).Description("URL of the dashboard.").
			Extractor(column_value_extractor.StructSelector("Url")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("layout_type").ColumnType(schema.ColumnTypeString).Description("Layout type of the dashboard. Can be on of \"free\" or \"ordered\".").
			Extractor(column_value_extractor.StructSelector("LayoutType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the dashboard.").
			Extractor(column_value_extractor.StructSelector("Title")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeTimestamp).Description("Creation date of the dashboard.").
			Extractor(column_value_extractor.StructSelector("CreatedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("author_handle").ColumnType(schema.ColumnTypeString).Description("Identifier of the dashboard author.").
			Extractor(column_value_extractor.StructSelector("AuthorHandle")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Description("Description of the dashboard.").
			Extractor(column_value_extractor.StructSelector("Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_read_only").ColumnType(schema.ColumnTypeBool).Description("Indicates if the dashboard is read-only. If True, only the author and admins can make changes to it.").
			Extractor(column_value_extractor.StructSelector("IsReadOnly")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("modified_at").ColumnType(schema.ColumnTypeTimestamp).Description("Modification time of the dashboard.").
			Extractor(column_value_extractor.StructSelector("ModifiedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("Dashboard identifier.").Build(),
	}
}

func (x *TableDatadogDashboardGenerator) GetSubTables() []*schema.Table {
	return nil
}
