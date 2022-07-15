package server

import (
	"net/http"
	"time"

	"github.com/yoshi429/test/model"
	"github.com/yoshi429/test/router"
)

func New() *http.Server {
	d := model.New()
	r := router.New(d)

	return &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}
