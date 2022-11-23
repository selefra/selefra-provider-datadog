package tables

import (
	"context"

	"github.com/DataDog/datadog-api-client-go/api/v2/datadog"
	"github.com/selefra/selefra-provider-datadog/datadog_client"
	"github.com/selefra/selefra-provider-datadog/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableDatadogLogEventGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableDatadogLogEventGenerator{}

func (x *TableDatadogLogEventGenerator) GetTableName() string {
	return "datadog_log_event"
}

func (x *TableDatadogLogEventGenerator) GetTableDescription() string {
	return ""
}

func (x *TableDatadogLogEventGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableDatadogLogEventGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableDatadogLogEventGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			ctx, apiClient, _, err := datadog_client.V2(ctx, taskClient.(*datadog_client.Client).Config)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			sort := datadog.LogsSort("timestamp")
			opts := *datadog.NewListLogsGetOptionalParameters()
			opts.WithSort(sort)
			opts.WithPageLimit(100)

			for {

				resp, _, err := apiClient.LogsApi.ListLogsGet(ctx, opts)
				if err != nil {

					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}

				for _, log := range resp.GetData() {
					resultChannel <- log

				}

				if resp.HasLinks() {
					if resp.Links.HasNext() {
						opts.WithPageCursor(*resp.Meta.GetPage().After)
					} else {
						break
					}
				} else {
					break
				}
			}

			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

func (x *TableDatadogLogEventGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableDatadogLogEventGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).Description("Status of the message associated with log.").
			Extractor(column_value_extractor.StructSelector("Attributes.Status")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("host").ColumnType(schema.ColumnTypeString).Description("Name of the machine from where the logs are being sent.").
			Extractor(column_value_extractor.StructSelector("Attributes.Host")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("message").ColumnType(schema.ColumnTypeString).Description("The message of the log.").
			Extractor(column_value_extractor.StructSelector("Attributes.Message")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("attributes").ColumnType(schema.ColumnTypeJSON).Description("JSON object of attributes for log.").
			Extractor(column_value_extractor.StructSelector("Attributes.Attributes")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Description("Array of tags associated with log.").
			Extractor(column_value_extractor.StructSelector("Attributes.Tags")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("Unique ID of the Log.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("timestamp").ColumnType(schema.ColumnTypeTimestamp).Description("Timestamp of log.").
			Extractor(column_value_extractor.StructSelector("Attributes.Timestamp")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("service").ColumnType(schema.ColumnTypeString).Description("The name of the application or service generating the log events.").
			Extractor(column_value_extractor.StructSelector("Attributes.Service")).Build(),
	}
}

func (x *TableDatadogLogEventGenerator) GetSubTables() []*schema.Table {
	return nil
}
