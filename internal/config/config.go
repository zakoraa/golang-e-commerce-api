package config

import "github.com/spf13/viper"

type Config struct {
	AppName string
	AppPort string
	DB DBConfig
	Redis RedisConfig
	JWT JWTConfig
}

type DBConfig struct {
	Host string
	Port string
	User string
	Password string
	Name string
}

type RedisConfig struct {
	Addr string
}

type JWTConfig struct {
	Secret string
}

func Load() (*Config, error) {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	return &Config{
		AppName: viper.GetString("APP_NAME"),
		AppPort: viper.GetString("APP_PORT"),
		DB: DBConfig{
			Host: viper.GetString("DB_HOST"),
			Port: viper.GetString("DB_PORT"),
			User: viper.GetString("DB_USER"),
			Password: viper.GetString("DB_PASSWORD"),
			Name: viper.GetString("DB_NAME"),
		},
		Redis: RedisConfig{
			Addr: viper.GetString("REDIS_ADOR"),
		},
		JWT: JWTConfig{
			Secret: viper.GetString("JWT_SECRET"),
		},

	},nil
}