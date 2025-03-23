package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	ServerPort string `mapstructure:"SERVER_PORT"`
	JWTSecret  string `mapstructure:"JWT_SECRET"`
	JWTIssuer  string `mapstructure:"JWT_ISSUER"`
	Database   struct {
		Driver string `mapstructure:"DB_DRIVER"`
		DSN    string `mapstructure:"DB_DSN"`
	} `mapstructure:"DATABASE"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
