package configs

import (
	"github.com/spf13/viper"
	"log"
	"time"
)

type Config struct {
	AppPort       string        `yaml:"appPort"`
	DBUser        string        `yaml:"dbUser"`
	DBPass        string        `yaml:"dbPass"`
	DBName        string        `yaml:"dbName"`
	DBPort        string        `yaml:"dbPort"`
	DBHost        string        `yaml:"dbHost"`
	JWTSigningKey string        `yaml:"jwtSigningKey"`
	AccessTTL     time.Duration `yaml:"accessTTL"`
	RefreshTTL    time.Duration `yaml:"refreshTTL"`
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
	config = &Config{}
	if err := viper.Unmarshal(config); err != nil {
		log.Fatal("unable to decode config in struct!")
	}
	return config
}
