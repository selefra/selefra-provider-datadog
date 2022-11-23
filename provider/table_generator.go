package provider

import (
	"github.com/selefra/selefra-provider-datadog/table_schema_generator"
	"github.com/selefra/selefra-provider-datadog/tables"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
)

func GenTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&tables.TableDatadogUserGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableDatadogDashboardGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableDatadogLogEventGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableDatadogMonitorGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableDatadogPermissionGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableDatadogRoleGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableDatadogIntegrationAwsGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableDatadogLogsMetricGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableDatadogSecurityMonitoringRuleGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableDatadogSecurityMonitoringSignalGenerator{}),
	}
}
