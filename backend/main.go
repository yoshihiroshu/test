package main

import (
	"log"

	_ "github.com/lib/pq"
	"github.com/yoshi429/test/config"
	"github.com/yoshi429/test/server"
)

func main() {
	conf := config.New()

	s := server.New(conf)

	log.Fatalln(s.ListenAndServe())
}
