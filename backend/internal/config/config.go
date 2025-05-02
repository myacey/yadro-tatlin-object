package config

import (
	"log"

	"yadro-tatlin-object/internal/pkg/web"
	"yadro-tatlin-object/internal/repository"

	"github.com/spf13/viper"
)

type Config struct {
	HTTPServer web.ServerConfig  `mapstructure:"httpserver"`
	Repository repository.Config `mapstructure:"repository"`
}

func LoadConfig(cfgPath string) (Config, error) {
	viper.SetConfigFile(cfgPath)

	if err := viper.ReadInConfig(); err != nil {
		return Config{}, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return Config{}, err
	}

	log.Printf("config: %+v", config)
	return config, nil
}
