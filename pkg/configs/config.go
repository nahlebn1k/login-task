package configs

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	App struct {
		Port string
	}
	DB struct {
		User string
		Pass string
		Name string
		Port string
		Host string
	}
}

func InitConfig() error {
	viper.AddConfigPath("pkg/configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

var config *Config

func GetConfig() *Config {
	if err := InitConfig(); err != nil {
		log.Fatal("cant init config file!")
	}
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal("unable to decode config in struct!")
	}
	return config
}
