package config

import (
	"errors"

	"github.com/spf13/viper"
)

type Config struct {
	AppName string
	AppPort string
	DB      DBConfig
	Redis   RedisConfig
	JWT     JWTConfig
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	URL      string
}

type RedisConfig struct {
	Addr string
}

type JWTConfig struct {
	Secret string
}

func Load() (*Config, error) {
	viper.AutomaticEnv()

	cfg := &Config{
		AppName: viper.GetString("APP_NAME"),
		AppPort: viper.GetString("APP_PORT"),

		DB: DBConfig{
			Host:     viper.GetString("DB_HOST"),
			Port:     viper.GetString("DB_PORT"),
			User:     viper.GetString("DB_USER"),
			Password: viper.GetString("DB_PASSWORD"),
			Name:     viper.GetString("DB_NAME"),
			URL:      viper.GetString("DB_URL"),
		},

		Redis: RedisConfig{
			Addr: viper.GetString("REDIS_ADDR"),
		},

		JWT: JWTConfig{
			Secret: viper.GetString("JWT_SECRET"),
		},
	}

	if cfg.AppPort == "" {
		cfg.AppPort = "8080"
	}

	if cfg.DB.URL == "" {
		return nil, errors.New("DB_URL is required")
	}

	if cfg.JWT.Secret == "" {
		return nil, errors.New("JWT_SECRET is required")
	}

	return cfg, nil
}
