# Table: datadog_integration_aws

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_specific_namespace_rules | json | X | √ | An object, that enables or disables metric collection for specific AWS namespaces for this AWS account only. | 
| excluded_regions | json | X | √ | An array of AWS regions to exclude from metrics collection. | 
| filter_tags | json | X | √ | List of tags (in the form 'key:value') that define a filter which is used when collecting EC2 or Lambda resources. These key:value pairs can be used to both whitelist and blacklist tags. | 
| host_tags | json | X | √ | Array of tags (in the form `key:value`) to add to all hosts and metrics reporting through this integration. | 
| account_id | string | X | √ | Your AWS Account ID without dashes. | 
| cspm_resource_collection_enabled | bool | X | √ | Whether Datadog collects cloud security posture management resources from your AWS account. This includes additional resources not covered under the general `resource_collection`. | 
| secret_access_key | string | X | √ | Your AWS secret access key. Only required if your AWS account is a GovCloud or China account. | 
| metrics_collection_enabled | bool | X | √ | Whether Datadog collects metrics for this AWS account. | 
| resource_collection_enabled | bool | X | √ | Whether Datadog collects a standard set of resources from your AWS account. | 
| role_name | string | X | √ | Your Datadog role delegation name. | 
| access_key_id | string | X | √ | Your AWS access key ID. Only required if your AWS account is a GovCloud or China account. | 


