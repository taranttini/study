package configs

import (
	"github.com/spf13/viper"
)

var cfg *conf

type conf struct {
	HttpTimeout   int    `mapstructure:"HTTP_TIMEOUT"`
	UrlBrasilApi  string `mapstructure:"URL_BRASILAPI"`
	UrlViaCep     string `mapstructure:"URL_VIACEP"`
	WebServerPort string `mapstructure:"WEB_SERVER_PORT"`
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
