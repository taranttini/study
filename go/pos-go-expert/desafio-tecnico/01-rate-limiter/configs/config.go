package configs

import (
	"github.com/spf13/viper"
)

type localConfig struct {
	QTY_REQUEST_IP              int `mapstructure:"QTY_REQUEST_IP"`
	QTY_REQUEST_TOKEN           int `mapstructure:"QTY_REQUEST_TOKEN"`
	KEEP_REQUEST_PER_X_SECONDS  int `mapstructure:"KEEP_REQUEST_PER_X_SECONDS"`
	BLOCKED_TOKEN_PER_X_SECONDS int `mapstructure:"BLOCKED_TOKEN_PER_X_SECONDS"`
	BLOCKED_IP_PER_X_SECONDS    int `mapstructure:"BLOCKED_IP_PER_X_SECONDS"`
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
