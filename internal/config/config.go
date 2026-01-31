package config

import "github.com/spf13/viper"

type Config struct {
	AppPort string
	DBUrl string
}

func Load() (*Config, error){
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	_ = viper.ReadInConfig()

	cfg := &Config{
		AppPort: viper.GetString("APP_PORT"),
		DBUrl: viper.GetString("DATABASE_URL"),
	}

	return cfg, nil
}