package web

import (
	"database/sql"
	"fmt"
	"github.com/p4vlowVl4d/web-panel/model"
	"net/http"
)

func CreatePanel(conf model.ServerConfig) Panel {
	addr := fmt.Sprintf("%s:%s", conf.Host, conf.Port)
	server := http.Server{
		Addr:              addr,
		Handler:           NewRouter(),
	}
	return Panel{Server: server}
}

type Panel struct {
	Server http.Server
	DB     sql.DB
}

func (p Panel) StartServer() error {
	err := p.Server.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}
