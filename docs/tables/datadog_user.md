# Table: datadog_user

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| handle | string | X | √ | Handle of the user. | 
| disabled | bool | X | √ | Indicates if the user is disabled. | 
| email | string | X | √ | Email of the user. | 
| name | string | X | √ | Name of the user. | 
| icon | string | X | √ | URL of the user's icon. | 
| verified | bool | X | √ | Indicates the verification status of the user. | 
| relationships | json | X | √ | Relationships of the user object returned by the API. |