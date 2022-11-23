# Table: datadog_role

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| modified_at | timestamp | X | √ | Time of last role modification. | 
| users | json | X | √ | Set of objects containing the permission ID and the name of the permissions granted to this role. | 
| permissions | json | X | √ | List of users emails attached to role. | 
| name | string | X | √ | Name of the role. | 
| id | string | X | √ | Id of the role. | 
| user_count | int | X | √ | Number of users associated with the role. | 
| created_at | timestamp | X | √ | Creation time of the role. | 


