package web

import (
	"fmt"
	"net/http"
	"github.com/go-chi/chi"
)

func NewRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/", index)
	return r
}

func index(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w, "New web Panel")
}
