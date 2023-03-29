package tables

import (
	"context"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
	"github.com/selefra/selefra-provider-datadog/datadog_client"
	"github.com/selefra/selefra-provider-datadog/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableDatadogUserGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableDatadogUserGenerator{}

func (x *TableDatadogUserGenerator) GetTableName() string {
	return "datadog_user"
}

func (x *TableDatadogUserGenerator) GetTableDescription() string {
	return ""
}

func (x *TableDatadogUserGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableDatadogUserGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableDatadogUserGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			ctx, apiClient, _, err := datadog_client.Server(ctx, taskClient.(*datadog_client.Client).Config)

			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			api := datadogV1.NewUsersApi(apiClient)
			resp, _, err := api.ListUsers(ctx)

			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			for _, user := range resp.GetUsers() {
				resultChannel <- user
			}

			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)
		},
	}
}

func (x *TableDatadogUserGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableDatadogUserGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("handle").ColumnType(schema.ColumnTypeString).Description("Handle of the user.").
			Extractor(column_value_extractor.StructSelector("Handle")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("disabled").ColumnType(schema.ColumnTypeBool).Description("Indicates if the user is disabled.").
			Extractor(column_value_extractor.StructSelector("Disabled")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("email").ColumnType(schema.ColumnTypeString).Description("Email of the user.").
			Extractor(column_value_extractor.StructSelector("Email")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("Name of the user.").
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		//table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("Id of the user.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("icon").ColumnType(schema.ColumnTypeString).Description("URL of the user's icon.").
			Extractor(column_value_extractor.StructSelector("Icon")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("verified").ColumnType(schema.ColumnTypeBool).Description("Indicates the verification status of the user.").
			Extractor(column_value_extractor.StructSelector("Verified")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("access_role").ColumnType(schema.ColumnTypeString).Description("The access role of the user. Options are **st** (standard user), **adm** (admin user), or **ro** (read-only user).").
			Extractor(column_value_extractor.StructSelector("AccessRole")).Build(),
	}
}

func (x *TableDatadogUserGenerator) GetSubTables() []*schema.Table {
	return nil
}
