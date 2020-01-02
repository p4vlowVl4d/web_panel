package main

import (
	"flag"
	"github.com/p4vlowVl4d/web-panel/model"
	"github.com/p4vlowVl4d/web-panel/web"
	"github.com/spf13/viper"
	"log"
)
var (
	CONFIG_FILE = "/"
)

func init() {
	flag.StringVar(&CONFIG_FILE, "c", CONFIG_FILE, "путь к файлу конфигурации")
}

func main() {
	log.Print("Starting application...")
	flag.Parse()
	log.Print("Parse configurations")
	config := createServerConfig()
	app := web.CreatePanel(config)
	log.Printf("Server listen on: %s:%s", config.Host, config.Port)
	err := app.StartServer()
	if err != nil {
		log.Fatal(err)
	}
}

func createServerConfig() model.ServerConfig {
	var config model.ServerConfig
	if CONFIG_FILE == "/" || CONFIG_FILE == "" {
		viper.SetConfigFile("config/app.yaml")
	} else {
		viper.SetConfigFile(CONFIG_FILE)
	}
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
	config.Host = viper.GetString("server.host")
	config.Port = viper.GetString("server.port")
	config.ReadTime = viper.GetDuration("server.readTimeout")
	config.WriteTime = viper.GetDuration("server.writeTimeout")
	return config
}