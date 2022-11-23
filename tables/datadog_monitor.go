package tables

import (
	"context"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"github.com/selefra/selefra-provider-datadog/datadog_client"
	"github.com/selefra/selefra-provider-datadog/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableDatadogMonitorGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableDatadogMonitorGenerator{}

func (x *TableDatadogMonitorGenerator) GetTableName() string {
	return "datadog_monitor"
}

func (x *TableDatadogMonitorGenerator) GetTableDescription() string {
	return ""
}

func (x *TableDatadogMonitorGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableDatadogMonitorGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableDatadogMonitorGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			ctx, apiClient, err := datadog_client.V1(ctx, taskClient.(*datadog_client.Client).Config)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			opts := datadog.ListMonitorsOptionalParameters{}

			resp, _, err := apiClient.MonitorsApi.ListMonitors(ctx, opts)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			for _, monitor := range resp {
				resultChannel <- monitor

			}

			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

func (x *TableDatadogMonitorGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableDatadogMonitorGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("query").ColumnType(schema.ColumnTypeString).Description("The monitor query.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("modified_at").ColumnType(schema.ColumnTypeTimestamp).Description("Last timestamp when the monitor was edited.").
			Extractor(column_value_extractor.StructSelector("Modified")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).Description("The type of the monitor. For more information about type, see https://docs.datadoghq.com/monitors/guide/monitor_api_options/.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("overall_state").ColumnType(schema.ColumnTypeString).Description("Current state of the monitor. Possible states are \"Alert\", \"Ignored\", \"No Data\", \"OK\", \"Skipped\", \"Unknown\" and \"Warn\".").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("options").ColumnType(schema.ColumnTypeJSON).Description("A list of role identifiers that can be pulled from the Roles API. Cannot be used with `locked` option.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("ID of the monitor.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("message").ColumnType(schema.ColumnTypeString).Description("Timestamp of the monitor creation.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("priority").ColumnType(schema.ColumnTypeInt).Description("Integer from 1 (high) to 5 (low) indicating alert severity.").
			Extractor(column_value_extractor.StructSelector("Priority")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("restricted_roles").ColumnType(schema.ColumnTypeJSON).Description("Relationships of the user object returned by the API.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Description("Tags associated to monitor.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("Name of the monitor.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeTimestamp).Description("Timestamp of the monitor creation.").
			Extractor(column_value_extractor.StructSelector("Created")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("multi").ColumnType(schema.ColumnTypeBool).Description("Whether or not the monitor is broken down on different groups.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("group_states").ColumnType(schema.ColumnTypeJSON).Description("Dictionary where the keys are groups (comma separated lists of tags) and the values are the list of groups your monitor is broken down on.").
			Extractor(column_value_extractor.StructSelector("State.Groups")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("creator_email").ColumnType(schema.ColumnTypeString).Description("Email of the creator.").
			Extractor(column_value_extractor.StructSelector("Creator.Email")).Build(),
	}
}

func (x *TableDatadogMonitorGenerator) GetSubTables() []*schema.Table {
	return nil
}
