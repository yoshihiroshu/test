package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yoshi429/test/handler"
	"github.com/yoshi429/test/request"
)

func ApplyRouters(r *mux.Router, rc *request.Context) {
	h := handler.Handler{
		Context: rc,
	}

	r.HandleFunc("/", h.Index).Methods(http.MethodGet)

	t := r.PathPrefix("/test").Subrouter()
	t.HandleFunc("", h.TestHandler).Methods(http.MethodGet)

	c := r.PathPrefix("/cmd").Subrouter()
	c.HandleFunc("", h.Command).Methods(http.MethodGet)
}
