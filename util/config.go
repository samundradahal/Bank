package util

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Dbdriver            string        `mapstructure:"DB_SERVER"`
	Serverport          string        `mapstructure:"SERVER_PORT"`
	TokenSymmetricKey   string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig() //Start reading config file

	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
