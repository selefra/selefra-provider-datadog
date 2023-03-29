package tables

import (
	"context"
	"github.com/DataDog/datadog-api-client-go/api/v2/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/selefra/selefra-provider-datadog/datadog_client"
	"github.com/selefra/selefra-provider-datadog/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableDatadogRoleGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableDatadogRoleGenerator{}

func (x *TableDatadogRoleGenerator) GetTableName() string {
	return "datadog_role"
}

func (x *TableDatadogRoleGenerator) GetTableDescription() string {
	return ""
}

func (x *TableDatadogRoleGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableDatadogRoleGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableDatadogRoleGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			ctx, apiClient, _, err := datadog_client.Server(ctx, taskClient.(*datadog_client.Client).Config)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			opts := datadogV2.ListRolesOptionalParameters{
				PageSize:   datadog.PtrInt64(int64(100)),
				PageNumber: datadog.PtrInt64(int64(0)),
			}

			count := int64(0)

			api := datadogV2.NewRolesApi(apiClient)

			for {
				resp, _, err := api.ListRoles(ctx, opts)

				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}

				for _, role := range resp.GetData() {
					count++
					resultChannel <- role
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
		},
	}
}

func userList(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (interface{}, error) {
	users := result.([]datadog.User)

	var user_emails []string

	for _, user := range users {
		user_emails = append(user_emails, *user.Attributes.Email)
	}

	return user_emails, nil
}

func (x *TableDatadogRoleGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableDatadogRoleGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("modified_at").ColumnType(schema.ColumnTypeTimestamp).Description("Time of last role modification.").
			Extractor(column_value_extractor.StructSelector("Attributes.ModifiedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("users").ColumnType(schema.ColumnTypeJSON).Description("Set of objects containing the permission ID and the name of the permissions granted to this role.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				r, err := userList(ctx, clientMeta, client, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				return r, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("permissions").ColumnType(schema.ColumnTypeJSON).Description("List of users emails attached to role.").
			Extractor(column_value_extractor.StructSelector("Relationships.Permissions.Data")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("Name of the role.").
			Extractor(column_value_extractor.StructSelector("Attributes.Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("Id of the role.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("user_count").ColumnType(schema.ColumnTypeInt).Description("Number of users associated with the role.").
			Extractor(column_value_extractor.StructSelector("Attributes.UserCount")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeTimestamp).Description("Creation time of the role.").
			Extractor(column_value_extractor.StructSelector("Attributes.CreatedAt")).Build(),
	}
}

func (x *TableDatadogRoleGenerator) GetSubTables() []*schema.Table {
	return nil
}
