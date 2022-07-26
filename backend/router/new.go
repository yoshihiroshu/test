package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yoshi429/test/config"
	"github.com/yoshi429/test/handler"
	"github.com/yoshi429/test/request"
)

func New(conf config.Configs) http.Handler {
	r := mux.NewRouter()

	rc := request.NewContext(conf)

	h := handler.Handler{
		Context: rc,
	}

	r.Use(rc.TestMiddleware)

	r.HandleFunc("/", h.Index).Methods(http.MethodGet)

	t := r.PathPrefix("/test").Subrouter()
	t.HandleFunc("", h.TestHandler).Methods(http.MethodGet)

	return r
}
