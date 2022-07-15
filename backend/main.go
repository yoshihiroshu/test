package main

import (
	"log"

	_ "github.com/lib/pq"
	"github.com/yoshi429/test/server"
)

func main() {

	s := server.New()

	log.Fatalln(s.ListenAndServe())
}
