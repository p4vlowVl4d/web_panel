package web

import (
	"github.com/p4vlowVl4d/web-panel/model"
	"net/http"
)

func CreatePanel(conf model.ServerConfig) Panel {
	//var addr string
	//addr = fmt.Sprintf("%s:%s", conf.Host, conf.Port)
	server := http.Server{
		Addr:              "127.0.0.1:3030",
		Handler:           NewRouter(),
		TLSConfig:         nil,
		ReadTimeout:       conf.ReadTime,
		ReadHeaderTimeout: 0,
		WriteTimeout:      conf.WriteTime,
		IdleTimeout:       0,
		MaxHeaderBytes:    0,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
	}
	return Panel{Server: server}
}

type Panel struct {
	Server http.Server
}

func (p Panel) StartServer() error {
	err := p.Server.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}

func (p *Panel) SetHandler(h http.Handler) {
	p.Server.Handler = h
}
