package main

import (
	"fmt"
	"github.com/p4vlowVl4d/web-panel/model"
	"github.com/p4vlowVl4d/web-panel/web"
	"log"
	"github.com/spf13/viper"
)

func main() {
	log.Print("Starting application...")
	log.Print("Parse configurations")
	config := createServerConfig()
	app := web.CreatePanel(config)
	log.Printf("Server listen on: %S:%S", config.Host, config.Port)
	fmt.Printf("Server listen on: %s:%s", config.Host,config.Port)
	log.Fatal(app.StartServer())
}

func createServerConfig() model.ServerConfig {
	var config model.ServerConfig
	//viper.AddConfigPath("../config")
	viper.SetConfigFile("../config/app.yaml")
	config.Host = viper.GetString("server.host")
	config.Host = viper.GetString("server.port")
	config.ReadTime = viper.GetDuration("server.readTimeout")
	config.WriteTime = viper.GetDuration("server.writeTimeout")
	return config
}