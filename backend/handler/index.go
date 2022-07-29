package handler

import (
	"fmt"
	"net/http"

	"github.com/yoshi429/test/request"
)

type Handler struct {
	Context *request.Context
}

func (h Handler) Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello - WORLD")
}

func (h Handler) TestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UNKOOOOO")
	fmt.Printf("RequestContext: %#v\n", h.Context)
	fmt.Fprintf(w, "RequestContext: %#v\n", h.Context)
	fmt.Println("UNKOOOOO")
}
