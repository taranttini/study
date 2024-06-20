package configs

import (
	"github.com/spf13/viper"
)

type conf struct {
	REQUEST_LIMIT_IP  string `mapstructure:"REQUEST_LIMIT_IP"`
	REQUEST_LIMIT_API string `mapstructure:"REQUEST_LIMIT_API"`
	TIME_LIMIT        string `mapstructure:"TIME_LIMIT"`
}

func LoadConfig(path string) (*conf, error) {
	var cfg *conf
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}
	return cfg, err
}
