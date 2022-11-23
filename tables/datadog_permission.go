package tables

import (
	"context"

	"github.com/selefra/selefra-provider-datadog/datadog_client"
	"github.com/selefra/selefra-provider-datadog/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableDatadogPermissionGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableDatadogPermissionGenerator{}

func (x *TableDatadogPermissionGenerator) GetTableName() string {
	return "datadog_permission"
}

func (x *TableDatadogPermissionGenerator) GetTableDescription() string {
	return ""
}

func (x *TableDatadogPermissionGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableDatadogPermissionGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableDatadogPermissionGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			ctx, apiClient, _, err := datadog_client.V2(ctx, taskClient.(*datadog_client.Client).Config)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			resp, _, err := apiClient.RolesApi.ListPermissions(ctx)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			for _, permission := range resp.GetData() {
				resultChannel <- permission

			}

			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

func (x *TableDatadogPermissionGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableDatadogPermissionGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Description("Description of the permission.").
			Extractor(column_value_extractor.StructSelector("Attributes.Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("display_name").ColumnType(schema.ColumnTypeString).Description("Displayed name for the permission.").
			Extractor(column_value_extractor.StructSelector("Attributes.DisplayName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("display_type").ColumnType(schema.ColumnTypeString).Description("Displayed type the permission.").
			Extractor(column_value_extractor.StructSelector("Attributes.DisplayType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("Name of the permission.").
			Extractor(column_value_extractor.StructSelector("Attributes.Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("Id of the permission.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("restricted").ColumnType(schema.ColumnTypeBool).Description("Whether or not the permission is restricted.").
			Extractor(column_value_extractor.StructSelector("Attributes.Restricted")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("group_name").ColumnType(schema.ColumnTypeString).Description("Name of the permission group.").
			Extractor(column_value_extractor.StructSelector("Attributes.GroupName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeTimestamp).Description("Creation time of the permission.").
			Extractor(column_value_extractor.StructSelector("Attributes.Created")).Build(),
	}
}

func (x *TableDatadogPermissionGenerator) GetSubTables() []*schema.Table {
	return nil
}
