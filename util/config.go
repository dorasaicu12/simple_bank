package util

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DBDriver             string        `mapstructure:"DB_DRIVER"`
	DBSource             string        `mapstructure:"DB_SOURCE"`
	ServerAddress        string        `mapstructure:"SERVER_ADDRESS"`
	GrpcServerAddress    string        `mapstructure:"GRPC_SERVER_ADDRESS"`
	TokenSymmectricKey   string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	TokenDuration        time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
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
	err = viper.Unmarshal(&config)
	if err != nil {
		return
	}
	// Trim spaces manually
	config.DBDriver = strings.TrimSpace(config.DBDriver)
	config.DBSource = strings.TrimSpace(config.DBSource)
	config.ServerAddress = strings.TrimSpace(config.ServerAddress)
	config.TokenSymmectricKey = strings.TrimSpace(config.TokenSymmectricKey)
	fmt.Println("DB Driver:", config.DBDriver)
	return
}
