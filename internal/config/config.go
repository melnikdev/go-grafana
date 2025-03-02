package config

import (
	"strings"
	"sync"

	"github.com/spf13/viper"
)

type Config struct {
	MongoDB *MongoDB
	Server  *Server
}

type Server struct {
	Port int
}

type MongoDB struct {
	Uri      string
	Database string
}

var once sync.Once
var configInstance *Config

func MustLoad() *Config {
	once.Do(func() {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(".")
		viper.AutomaticEnv()
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

		if err := viper.ReadInConfig(); err != nil {
			panic(err)
		}

		if err := viper.Unmarshal(&configInstance); err != nil {
			panic(err)
		}
	})

	return configInstance
}
