package tables

import (
	"context"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/selefra/selefra-provider-datadog/datadog_client"
	"github.com/selefra/selefra-provider-datadog/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableDatadogLogsMetricGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableDatadogLogsMetricGenerator{}

func (x *TableDatadogLogsMetricGenerator) GetTableName() string {
	return "datadog_logs_metric"
}

func (x *TableDatadogLogsMetricGenerator) GetTableDescription() string {
	return ""
}

func (x *TableDatadogLogsMetricGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableDatadogLogsMetricGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableDatadogLogsMetricGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			ctx, apiClient, _, err := datadog_client.Server(ctx, taskClient.(*datadog_client.Client).Config)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			api := datadogV2.NewLogsMetricsApi(apiClient)
			resp, _, err := api.ListLogsMetrics(ctx)

			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			for _, logMetric := range resp.GetData() {
				resultChannel <- logMetric
			}

			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)
		},
	}
}

func (x *TableDatadogLogsMetricGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableDatadogLogsMetricGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("compute_path").ColumnType(schema.ColumnTypeString).Description("The path to the value the log-based metric will aggregate on (only used if the aggregation type is a \"distribution\").").
			Extractor(column_value_extractor.StructSelector("Attributes.Compute.Path")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("filter_query").ColumnType(schema.ColumnTypeString).Description("The search query - following the log search syntax to filter logs.").
			Extractor(column_value_extractor.StructSelector("Attributes.Filter.Query")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("group_by").ColumnType(schema.ColumnTypeJSON).Description("List of rules for the group by.").
			Extractor(column_value_extractor.StructSelector("Attributes.GroupBy")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("The name of the log-based metric.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("compute_aggregation_type").ColumnType(schema.ColumnTypeString).Description("The type of aggregation to used for computing metric. Can be one of \"count\", \"distribution\".").
			Extractor(column_value_extractor.StructSelector("Attributes.Compute.AggregationType")).Build(),
	}
}

func (x *TableDatadogLogsMetricGenerator) GetSubTables() []*schema.Table {
	return nil
}
