package model

import (
	"time"
)

type ServerConfig struct {
	Host, Port 			string
	ReadTime, WriteTime time.Duration
}
