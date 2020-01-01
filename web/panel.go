package web

import (
	"fmt"
	"github.com/p4vlowVl4d/web-panel/model"
	"log"
	"net/http"
)

func CreatePanel(conf model.ServerConfig) Panel {
	var addr string
	addr = fmt.Sprintf("%s:%s", conf.Host, conf.Port)
	server := http.Server{
		Addr:              addr,
		Handler:           nil,
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

func (p Panel) StartServer() {
	log.Fatal(p.Server.ListenAndServe())
}

func (p *Panel) SetHandler(h http.Handler) {

}
