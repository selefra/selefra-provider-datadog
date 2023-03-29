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
		//table_schema_generator.GenTableSchema(&tables.TableDatadogLogEventGenerator{}),
		//table_schema_generator.GenTableSchema(&tables.TableDatadogMonitorGenerator{}),xx
		//table_schema_generator.GenTableSchema(&tables.TableDatadogPermissionGenerator{}), xx
		//table_schema_generator.GenTableSchema(&tables.TableDatadogRoleGenerator{}),xx
		//table_schema_generator.GenTableSchema(&tables.TableDatadogIntegrationAwsGenerator{}),xx
		//table_schema_generator.GenTableSchema(&tables.TableDatadogLogsMetricGenerator{}), xx
		//table_schema_generator.GenTableSchema(&tables.TableDatadogSecurityMonitoringRuleGenerator{}),xx
		//table_schema_generator.GenTableSchema(&tables.TableDatadogSecurityMonitoringSignalGenerator{}),xx
	}
}
