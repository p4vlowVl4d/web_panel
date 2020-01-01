package main

import (
	"github.com/p4vlowVl4d/web-panel/model"
	"log"
	"github.com/spf13/viper"
)

func main() {
	log.Print("Starting application")
	viper.AddConfigPath("./config")
	viper.SetConfigFile("app.yaml")
}

func createServerConfig() model.ServerConfig {
	var config model.ServerConfig
	config.Host = viper.GetString("server.host")
	config.Host = viper.GetString("server.port")
	config.ReadTime = viper.GetDuration("server.readTimeout")
	config.WriteTime = viper.GetDuration("server.writeTimeout")
	return config
}