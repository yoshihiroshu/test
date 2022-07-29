package request

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/yoshi429/test/cache"
	"github.com/yoshi429/test/config"
	"github.com/yoshi429/test/model"
)

type Context struct {
	Db     *model.DBContext
	Cache  *cache.RedisContext
	Conf   config.Configs
	Logger *log.Logger
}

func NewContext(conf config.Configs) *Context {
	return &Context{
		Db:     model.New(conf),
		Cache:  cache.New(conf),
		Conf:   conf,
		Logger: log.New(os.Stdout, "", log.LstdFlags),
	}
}

// 廃止
func (c Context) Handler(next func(http.ResponseWriter, *http.Request, Context)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		next(w, r, c)
	}
}

func (c Context) TestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c.Logger.Println("url:", r.URL)
		fmt.Fprintln(w, "url:", r.URL)
		next.ServeHTTP(w, r)
	})
}
