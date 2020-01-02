package web

import (
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
)

func NewRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/", index)
	r.Get("/login", auth)
	r.Handle("/assets/*",
		http.StripPrefix("/assets/",
			http.FileServer(http.Dir("public/"))))
	return r
}

func index(w http.ResponseWriter, r *http.Request)  {
	tmpl, _ := template.ParseFiles("resources/view/index.html")
	tmpl.Execute(w, nil)
	//fmt.Fprintf(w, "New web Panel.\nYou Ip: %s", r.RemoteAddr)
}

func auth(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("resources/view/security/login.html")
	tmpl.Execute(w, nil)
}
