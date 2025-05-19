package util

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DBDriver           string `mapstructure:"DB_DRIVER"`
	DBSource           string `mapstructure:"DB_SOURCE"`
	ServerAddress      string `mapstructure:"SERVER_ADDRESS"`
	TokenSymmectricKey string `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	TokenDuration      time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal("err config evn", err.Error())
		return
	}
	viper.Unmarshal(&config)
	return
}
