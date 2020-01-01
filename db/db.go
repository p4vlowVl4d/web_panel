package db

import (
	"github.com/jmoiron/sqlx"
)

type DB struct {
	Conn sqlx.DB
}

type DBConfig struct {
	Host, Port string
	User, Password string
	DbName string
}

func getConnection(config DBConfig) DB {
	conn := sqlx.DB{
		DB:     nil,
		Mapper: nil,
	}
	return DB{}
}


