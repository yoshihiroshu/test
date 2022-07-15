package request

import (
	"net/http"

	"github.com/yoshi429/test/model"
)

type Context struct {
	Db *model.DBContext
}

func NewContext(db *model.DBContext) *Context {
	return &Context{
		Db: db,
	}
}

func (rc Context) Handler(next func(http.ResponseWriter, *http.Request, Context)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		next(w, r, rc)
	}
}
