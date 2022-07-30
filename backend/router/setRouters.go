package router

import (
	"net/http"

	"github.com/yoshi429/test/handler"
)

func (r Router) ApplyRouters() {
	r.Use(r.Context.TestMiddleware)

	h := handler.Handler{
		Context: r.Context,
	}

	r.Handle("/", AppHandler(h.Index)).Methods(http.MethodGet)

	t := r.PathPrefix("/test").Subrouter()
	t.Handle("", AppHandler(h.TestHandler)).Methods(http.MethodGet)

	c := r.PathPrefix("/cmd").Subrouter()
	c.Handle("", AppHandler(h.Command)).Methods(http.MethodGet)

	/*
		r.AppHandle("/", h.Index).Methods(http.MethodGet)

		// Grouping
		t := r.Group("/test")
		t.AppHandle("", h.TestHandler).Methods(http.MethodGet)

		c := r.Group("/cmd")
		c.AppHandle("", h.Command).Methods(http.MethodGet)

		user := r.Group("/user")
		user.HandleFunc("/register", h.RegisterAccount).Methods(http.MethodPost)
	*/
}
