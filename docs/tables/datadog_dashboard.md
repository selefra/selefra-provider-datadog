# Table: datadog_dashboard

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| url | string | X | √ | URL of the dashboard. | 
| widgets | json | X | √ | List of widgets to display on the dashboard. | 
| layout_type | string | X | √ | Layout type of the dashboard. Can be on of "free" or "ordered". | 
| reflow_type | string | X | √ | Reflow type for a new dashboard layout dashboard. If set to 'fixed', the dashboard expects all widgets to have a layout, and if it's set to 'auto', widgets should not have layouts. | 
| title | string | X | √ | Title of the dashboard. | 
| template_variable_presets | json | X | √ | List of template variables saved views. | 
| created_at | timestamp | X | √ | Creation date of the dashboard. | 
| author_handle | string | X | √ | Identifier of the dashboard author. | 
| description | string | X | √ | Description of the dashboard. | 
| is_read_only | bool | X | √ | Indicates if the dashboard is read-only. If True, only the author and admins can make changes to it. | 
| modified_at | timestamp | X | √ | Modification time of the dashboard. | 
| restricted_roles | json | X | √ | A list of role identifiers. Only the author and users associated with at least one of these roles can edit this dashboard. Overrides the `is_read_only` property if both are present. | 
| template_variables | json | X | √ | List of template variables for this dashboard. | 
| id | string | X | √ | Dashboard identifier. | 


