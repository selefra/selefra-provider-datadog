# Table: datadog_monitor

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| query | string | X | √ | The monitor query. | 
| modified_at | timestamp | X | √ | Last timestamp when the monitor was edited. | 
| type | string | X | √ | The type of the monitor. For more information about type, see https://docs.datadoghq.com/monitors/guide/monitor_api_options/. | 
| overall_state | string | X | √ | Current state of the monitor. Possible states are "Alert", "Ignored", "No Data", "OK", "Skipped", "Unknown" and "Warn". | 
| options | json | X | √ | A list of role identifiers that can be pulled from the Roles API. Cannot be used with `locked` option. | 
| id | string | X | √ | ID of the monitor. | 
| message | string | X | √ | Timestamp of the monitor creation. | 
| priority | int | X | √ | Integer from 1 (high) to 5 (low) indicating alert severity. | 
| restricted_roles | json | X | √ | Relationships of the user object returned by the API. | 
| tags | json | X | √ | Tags associated to monitor. | 
| name | string | X | √ | Name of the monitor. | 
| created_at | timestamp | X | √ | Timestamp of the monitor creation. | 
| multi | bool | X | √ | Whether or not the monitor is broken down on different groups. | 
| group_states | json | X | √ | Dictionary where the keys are groups (comma separated lists of tags) and the values are the list of groups your monitor is broken down on. | 
| creator_email | string | X | √ | Email of the creator. | 


