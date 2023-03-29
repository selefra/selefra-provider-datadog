# Table: datadog_dashboard

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| url | string | X | √ | URL of the dashboard. | 
| layout_type | string | X | √ | Layout type of the dashboard. Can be on of "free" or "ordered". | 
| title | string | X | √ | Title of the dashboard. | 
| created_at | timestamp | X | √ | Creation date of the dashboard. | 
| author_handle | string | X | √ | Identifier of the dashboard author. | 
| description | string | X | √ | Description of the dashboard. | 
| is_read_only | bool | X | √ | Indicates if the dashboard is read-only. If True, only the author and admins can make changes to it. | 
| modified_at | timestamp | X | √ | Modification time of the dashboard. | 
| id | string | X | √ | Dashboard identifier. |