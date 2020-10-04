package config

import (
	"github.com/spf13/viper"
	"log"
)

var Config *viper.Viper

func SetConfig(filePath string) {
	log.Printf("[config] run the env with:%s", filePath)

	Config = viper.New()
	Config.SetConfigFile(filePath)
	if err := Config.ReadInConfig(); err != nil {
		log.Fatalf("[config] read config err: %v", err)
	}

}
