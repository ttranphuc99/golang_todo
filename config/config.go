package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	SecretKeyJwt              string `mapstructure:"SECRET_KEY_JWT"`
	TokenCurrentUserId        string `mapstructure:"TOKEN_CURRENT_USER_ID"`
	TokenCurrentUserRole      string `mapstructure:"TOKEN_CURRENT_USER_ROLE"`
	TokenExp                  string `mapstructure:"TOKEN_EXP"`
	ConnectionStr             string `mapstructure:"CONNECTION_STR"`
	InvalidTodoStatusArgument string `mapstructure:"INVALID_TODO_STATUS_ARGUMENT"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		log.Println("error while reading config " + err.Error())
		return
	}

	err = viper.Unmarshal(&config)
	return
}
