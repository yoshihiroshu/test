package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yoshi429/test/config"
	"github.com/yoshi429/test/handler"
	"github.com/yoshi429/test/request"
)

type Router struct {
	*mux.Router
	Context    *request.Context
	AppHandler func(w http.ResponseWriter, r *http.Request) error
}

func New(conf config.Configs) *Router {
	return &Router{
		Router:  mux.NewRouter(),
		Context: request.NewContext(conf),
	}
}

func (r Router) ApplyRouters() {
	rc := r.Context

	r.Use(rc.TestMiddleware)

	h := handler.Handler{
		Context: rc,
	}

	r.HandleFunc("/", h.Index).Methods(http.MethodGet)

	t := r.PathPrefix("/test").Subrouter()
	t.HandleFunc("", h.TestHandler).Methods(http.MethodGet)

	c := r.PathPrefix("/cmd").Subrouter()
	c.HandleFunc("", h.Command).Methods(http.MethodGet)

	user := r.PathPrefix("/user").Subrouter()
	user.HandleFunc("/register", h.RegisterAccount).Methods(http.MethodPost)
}
