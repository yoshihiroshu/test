package request

import (
	"net/http"

	"github.com/yoshi429/test/config"
	"github.com/yoshi429/test/model"
)

type Context struct {
	Db   *model.DBContext
	conf config.Configs
}

func NewContext(db *model.DBContext) *Context {
	return &Context{
		Db:   db,
		conf: config.New(),
	}
}

func (c Context) GetConfig() config.Configs {
	return c.conf
}

func (rc Context) Handler(next func(http.ResponseWriter, *http.Request, Context)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		next(w, r, rc)
	}
}
