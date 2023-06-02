package config

import "github.com/spf13/viper"

type Config struct {
	AppPort        string `mapstructure:"APP_PORT"`
	DatabaseURL    string `mapstructure:"DATABASE_URL"`
	DatabaseDriver string `mapstructure:"DATABASE_DRIVER"`
	LogLevel       string `mapstructure:"LOG_LEVEL"`
}

func LoadConfig() (Config, error) {
	var config Config

	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return config, err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return config, err
	}

	return config, nil
}
