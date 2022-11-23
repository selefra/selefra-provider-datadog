# Table: datadog_logs_metric

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| compute_path | string | X | √ | The path to the value the log-based metric will aggregate on (only used if the aggregation type is a "distribution"). | 
| filter_query | string | X | √ | The search query - following the log search syntax to filter logs. | 
| group_by | json | X | √ | List of rules for the group by. | 
| id | string | X | √ | The name of the log-based metric. | 
| compute_aggregation_type | string | X | √ | The type of aggregation to used for computing metric. Can be one of "count", "distribution". | 


