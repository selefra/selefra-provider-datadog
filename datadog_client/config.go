package datadog_client

type Config struct {
	ApiKey string `yaml:"api_key,omitempty" mapstructure:"api_key"`
	AppKey string `yaml:"app_key,omitempty" mapstructure:"app_key"`
	ApiUrl string `yaml:"api_url,omitempty" mapstructure:"api_url"`
}
