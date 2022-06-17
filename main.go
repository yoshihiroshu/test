package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {

	m := http.NewServeMux()

	m.HandleFunc("/", index)

	s := &http.Server{
		Addr:           ":8080",
		Handler:        m,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatalln(s.ListenAndServe())
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello - Tests")
}
