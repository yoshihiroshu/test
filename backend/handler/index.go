package handler

import (
	"fmt"
	"net/http"

	"github.com/yoshi429/test/request"
)

type IndexHandler struct{}

func (h IndexHandler) Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello - WORLD")
}

func (h IndexHandler) TestHandler(w http.ResponseWriter, r *http.Request, rc request.Context) {
	fmt.Println("UNKOOOOO")
	fmt.Printf("RequestContext: %#v\n", rc)
	fmt.Fprintf(w, "RequestContext: %#v\n", rc)
	fmt.Println("UNKOOOOO")
}
