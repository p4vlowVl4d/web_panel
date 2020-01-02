package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"log"
	"fmt"
)

type DB struct {
	Conn *sql.DB
}

type ManagerInterface interface {
	find(identifier int)
}

type EntityInterface interface {
	getId() int
	getScheme() string
	getTableName() string
}

type DBConfig struct {
	Host, Port string
	User, Password string
	DbName string
}


func getConnection(config DBConfig) DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.DbName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	return DB{Conn: db}
}

func CreateConfig() DBConfig {
	viper.SetConfigFile("config/app.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	return DBConfig{
		Host:     viper.GetString("mysql.host"),
		Port:     viper.GetString("mysql.port"),
		User:     viper.GetString("mysql.user"),
		Password: viper.GetString("mysql.password"),
		DbName:   viper.GetString("mysql.dbname"),
	}
}