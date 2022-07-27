package server

import (
	"net/http"
	"time"

	"github.com/yoshi429/test/config"
	"github.com/yoshi429/test/router"
)

func New(conf config.Configs) *http.Server {
	r := router.New(conf)

	return &http.Server{
		Addr:           conf.GetUserAddr(),
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}
