package tables

import (
	"context"

	"github.com/DataDog/datadog-api-client-go/api/v2/datadog"
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

			ctx, apiClient, _, err := datadog_client.V2(ctx, taskClient.(*datadog_client.Client).Config)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			opts := datadog.ListUsersOptionalParameters{
				PageSize:   datadog.PtrInt64(int64(100)),
				PageNumber: datadog.PtrInt64(int64(0)),
				Sort:       datadog.PtrString("name"),
			}

			paging := true
			count := int64(0)

			for paging {
				resp, _, err := apiClient.UsersApi.ListUsers(ctx, opts)
				if err != nil {

					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}

				for _, user := range resp.GetData() {
					count++
					resultChannel <- user

				}

				if resp.Meta.Page.HasTotalFilteredCount() {
					if count >= resp.Meta.Page.GetTotalFilteredCount() {
						return schema.NewDiagnosticsErrorPullTable(task.Table, nil)
					}
				}

				if count >= resp.Meta.Page.GetTotalCount() {
					return schema.NewDiagnosticsErrorPullTable(task.Table, nil)
				}
				opts.WithPageNumber(*opts.PageNumber + 1)
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
			Extractor(column_value_extractor.StructSelector("Attributes.Handle")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("disabled").ColumnType(schema.ColumnTypeBool).Description("Indicates if the user is disabled.").
			Extractor(column_value_extractor.StructSelector("Attributes.Disabled")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the user.").
			Extractor(column_value_extractor.StructSelector("Attributes.Title")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("modified_at").ColumnType(schema.ColumnTypeTimestamp).Description("Time that the user was last modified.").
			Extractor(column_value_extractor.StructSelector("Attributes.ModifiedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("service_account").ColumnType(schema.ColumnTypeBool).Description("Indicates if the user is a service account.").
			Extractor(column_value_extractor.StructSelector("Attributes.ServiceAccount")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("role_ids").ColumnType(schema.ColumnTypeJSON).Description("A list of role IDs attached to user.").
			Extractor(column_value_extractor.StructSelector("Relationships.Roles.Data")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("email").ColumnType(schema.ColumnTypeString).Description("Email of the user.").
			Extractor(column_value_extractor.StructSelector("Attributes.Email")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("Name of the user.").
			Extractor(column_value_extractor.StructSelector("Attributes.Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeTimestamp).Description("Creation time of the user.").
			Extractor(column_value_extractor.StructSelector("Attributes.CreatedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("Id of the user.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("icon").ColumnType(schema.ColumnTypeString).Description("URL of the user's icon.").
			Extractor(column_value_extractor.StructSelector("Attributes.Icon")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).Description("Status of the user.").
			Extractor(column_value_extractor.StructSelector("Attributes.Status")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("verified").ColumnType(schema.ColumnTypeBool).Description("Indicates the verification status of the user.").
			Extractor(column_value_extractor.StructSelector("Attributes.Verified")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("relationships").ColumnType(schema.ColumnTypeJSON).Description("Relationships of the user object returned by the API.").Build(),
	}
}

func (x *TableDatadogUserGenerator) GetSubTables() []*schema.Table {
	return nil
}
