# Table: datadog_user

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| handle | string | X | √ | Handle of the user. | 
| disabled | bool | X | √ | Indicates if the user is disabled. | 
| title | string | X | √ | Title of the user. | 
| modified_at | timestamp | X | √ | Time that the user was last modified. | 
| service_account | bool | X | √ | Indicates if the user is a service account. | 
| role_ids | json | X | √ | A list of role IDs attached to user. | 
| email | string | X | √ | Email of the user. | 
| name | string | X | √ | Name of the user. | 
| created_at | timestamp | X | √ | Creation time of the user. | 
| id | string | X | √ | Id of the user. | 
| icon | string | X | √ | URL of the user's icon. | 
| status | string | X | √ | Status of the user. | 
| verified | bool | X | √ | Indicates the verification status of the user. | 
| relationships | json | X | √ | Relationships of the user object returned by the API. | 


