package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	// m := http.NewServeMux()

	r.HandleFunc("/", index)

	s := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatalln(s.ListenAndServe())
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello - Tests")
}
