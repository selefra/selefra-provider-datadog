package tables

import (
	"context"
	"time"

	"github.com/DataDog/datadog-api-client-go/api/v2/datadog"
	"github.com/selefra/selefra-provider-datadog/datadog_client"
	"github.com/selefra/selefra-provider-datadog/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableDatadogSecurityMonitoringSignalGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableDatadogSecurityMonitoringSignalGenerator{}

func (x *TableDatadogSecurityMonitoringSignalGenerator) GetTableName() string {
	return "datadog_security_monitoring_signal"
}

func (x *TableDatadogSecurityMonitoringSignalGenerator) GetTableDescription() string {
	return ""
}

func (x *TableDatadogSecurityMonitoringSignalGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableDatadogSecurityMonitoringSignalGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableDatadogSecurityMonitoringSignalGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			ctx, _, configuration, err := datadog_client.V2(ctx, taskClient.(*datadog_client.Client).Config)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			filterFrom := time.Now().AddDate(0, 0, -1)
			filterTo := time.Now()
			pageLimit := int32(50)

			opts := datadog.ListSecurityMonitoringSignalsOptionalParameters{
				FilterFrom: &filterFrom,
				FilterTo:   &filterTo,
				PageLimit:  &pageLimit,
			}

			opts.WithSort(datadog.SECURITYMONITORINGSIGNALSSORT_TIMESTAMP_ASCENDING)

			if opts.FilterTo == nil {
				opts.WithFilterTo(time.Now())
			}

			configuration.SetUnstableOperationEnabled("ListSecurityMonitoringSignals", true)
			apiClient := datadog.NewAPIClient(configuration)

			for {
				resp, _, err := apiClient.SecurityMonitoringApi.ListSecurityMonitoringSignals(ctx, opts)
				if err != nil {

					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}

				for _, securityMonitoringSignal := range resp.GetData() {
					resultChannel <- securityMonitoringSignal

				}

				if meta, ok := resp.GetMetaOk(); ok {
					if page, pageOk := meta.GetPageOk(); pageOk {
						if page.HasAfter() {
							opts.WithPageCursor(page.GetAfter())
						}
					}
				} else {
					break
				}
			}
			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

func (x *TableDatadogSecurityMonitoringSignalGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableDatadogSecurityMonitoringSignalGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the security signal").
			Extractor(column_value_extractor.StructSelector("Attributes.Attributes")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("attributes").ColumnType(schema.ColumnTypeJSON).Description("A JSON object of attributes in the security signal.").
			Extractor(column_value_extractor.StructSelector("Attributes.Attributes")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Description("An array of tags associated with the security signal.").
			Extractor(column_value_extractor.StructSelector("Attributes.Tags")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("The unique ID of the security signal.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("message").ColumnType(schema.ColumnTypeString).Description("The message in the security signal defined by the rule that generated the signal.").
			Extractor(column_value_extractor.StructSelector("Attributes.Message")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("timestamp").ColumnType(schema.ColumnTypeTimestamp).Description("The timestamp of the security signal.").
			Extractor(column_value_extractor.StructSelector("Attributes.Timestamp")).Build(),
	}
}

func (x *TableDatadogSecurityMonitoringSignalGenerator) GetSubTables() []*schema.Table {
	return nil
}
