package datadog_client

type Configs struct {
	Providers []Config `yaml:"providers"  mapstructure:"providers"`
}

type Config struct {
	ApiKey string `yaml:"api_key,omitempty" mapstructure:"api_key"`
	AppKey string `yaml:"app_key,omitempty" mapstructure:"app_key"`
	ApiUrl string `yaml:"api_url,omitempty" mapstructure:"api_url"`
}
