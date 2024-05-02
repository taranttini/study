package configs

import (
	_ "time"

	"github.com/spf13/viper"
)

var cfg *conf

type conf struct {
	HttpTimeout  int32  `mapstructure:"HTTP_TIMEOUT"`
	UrlBrasilApi string `mapstructure:"URL_BRASILAPI"`
	UrlViaCep    string `mapstructure:"URL_VIACEP"`
}

func LoadConfig(path string) (*conf, error) {
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
