# Table: datadog_log_event

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| status | string | X | √ | Status of the message associated with log. | 
| host | string | X | √ | Name of the machine from where the logs are being sent. | 
| message | string | X | √ | The message of the log. | 
| attributes | json | X | √ | JSON object of attributes for log. | 
| tags | json | X | √ | Array of tags associated with log. | 
| id | string | X | √ | Unique ID of the Log. | 
| timestamp | timestamp | X | √ | Timestamp of log. | 
| service | string | X | √ | The name of the application or service generating the log events. | 


