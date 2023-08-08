package config

import (
	"github.com/spf13/viper"
)

type ServerConfig struct {
	Port        int    `mapstructure:"PORT"`
	Environment string `mapstructure:"ENVIRONMENT"`
	Debug       bool   `mapstructure:"DEBUG"`
}

type DatabaseConfig struct {
	Host     string `mapstructure:"HOST"`
	Port     int    `mapstructure:"PORT"`
	Name     string `mapstructure:"NAME"`
	User     string `mapstructure:"USER"`
	Password string `mapstructure:"PASSWORD"`
	SSLMode  string `mapstructure:"SSLMODE"`
	FileName string `mapstructure:"FILENAME"`
}

type StripeConfig struct {
	Key string `mapstructure:"KEY"`
}

type Config struct {
	Server   ServerConfig   `mapstructure:"SERVER"`
	Database DatabaseConfig `mapstructure:"DATABASE"`
	Stripe   StripeConfig   `mapstructure:"STRIPE"`
}

func InitializeConfig() (*Config, error) {
	viper.SetConfigName(".env") // allow directly reading from .env file
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	viper.AddConfigPath("/")
	viper.AllowEmptyEnv(false)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	viper.AutomaticEnv()

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
