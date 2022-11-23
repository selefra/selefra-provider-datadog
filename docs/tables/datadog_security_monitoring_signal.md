# Table: datadog_security_monitoring_signal

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| title | string | X | √ | Title of the security signal | 
| attributes | json | X | √ | A JSON object of attributes in the security signal. | 
| tags | json | X | √ | An array of tags associated with the security signal. | 
| id | string | X | √ | The unique ID of the security signal. | 
| message | string | X | √ | The message in the security signal defined by the rule that generated the signal. | 
| timestamp | timestamp | X | √ | The timestamp of the security signal. | 


