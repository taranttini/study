package configs

import (
	"github.com/spf13/viper"
)

type localConfig struct {
	QTY_REQUEST_IP    int    `mapstructure:"QTY_REQUEST_IP"`
	QTY_REQUEST_TOKEN int    `mapstructure:"QTY_REQUEST_TOKEN"`
	TIME_LIMIT        string `mapstructure:"TIME_LIMIT"`
}

func LoadConfig(path string) (*localConfig, error) {
	var cfg *localConfig
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
