package request

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/yoshi429/test/config"
	"github.com/yoshi429/test/model"
)

type Context struct {
	Db     *model.DBContext
	conf   config.Configs
	Logger *log.Logger
}

func NewContext(conf config.Configs) *Context {
	return &Context{
		Db:     model.New(conf),
		conf:   conf,
		Logger: log.New(os.Stdout, "", log.LstdFlags),
	}
}

func (c Context) NewLogger(w io.Writer, s string, i int) {
	c.Logger = log.New(w, s, i)
}

func (c Context) GetConfig() config.Configs {
	return c.conf
}

func (c Context) Handler(next func(http.ResponseWriter, *http.Request, Context)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		next(w, r, c)
	}
}

func (c Context) TestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c.Logger.Println("url:", r.URL)
		next.ServeHTTP(w, r)
	})
}
