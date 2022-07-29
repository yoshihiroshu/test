package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yoshi429/test/config"
	"github.com/yoshi429/test/request"
)

func New(conf config.Configs) http.Handler {
	r := mux.NewRouter()

	rc := request.NewContext(conf)

	r.Use(rc.TestMiddleware)

	ApplyRouters(r, rc)

	return r
}
