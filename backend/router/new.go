package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yoshi429/test/handler"
	"github.com/yoshi429/test/model"
	"github.com/yoshi429/test/request"
)

func New(db *model.DBContext) http.Handler {
	r := mux.NewRouter()

	rc := request.NewContext(db)

	ih := &handler.IndexHandler{}

	r.HandleFunc("/", ih.Index)

	r.HandleFunc("/test", rc.Handler(ih.TestHandler))

	return r
}
