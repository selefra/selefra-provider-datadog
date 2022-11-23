# Table: datadog_security_monitoring_rule

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| message | string | X | √ | Message for generated signals. | 
| options | json | X | √ | Additional options for security monitoring rules. | 
| id | string | X | √ | The ID of the rule. | 
| update_author_id | string | X | √ | User ID of the user who updated the rule. | 
| created_at | string | X | √ | When the rule was created, timestamp in milliseconds. | 
| is_default | bool | X | √ | Whether the rule is included by default. | 
| version | int | X | √ | The version of the rule. | 
| tags | json | X | √ | Tags for generated signals. | 
| is_enabled | bool | X | √ | Whether the rule is enabled. | 
| queries | string | X | √ | Queries for selecting logs which are part of the rule. | 
| type | string | X | √ | The security monitoring rule type. | 
| cases | json | X | √ | Cases for generating signals. | 
| name | string | X | √ | The name of the rule. | 
| creation_author_id | string | X | √ | User ID of the user who created the rule. | 
| has_extended_title | bool | X | √ | Whether the notifications include the triggering group-by values in their title. | 
| is_deleted | bool | X | √ | Whether the rule has been deleted. | 
| filters | json | X | √ | Additional queries to filter matched events before they are processed. | 


