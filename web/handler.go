package web

import (
	"github.com/go-chi/chi"
	"html/template"
	"log"
	"net/http"
)

func NewRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/", index)
	r.Get("/login", authForm)
	r.Handle("/assets/*",
		http.StripPrefix("/assets/",
			http.FileServer(http.Dir("public/"))))
	return r
}

func index(w http.ResponseWriter, r *http.Request)  {
	tmpl, _ := template.ParseFiles("resources/view/index.html")
	log.Fatal(tmpl.Execute(w, nil))
}

func authForm(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("resources/view/security/login.html")
	log.Fatal(tmpl.Execute(w, nil))
}
