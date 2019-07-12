package http

import (
	l "github.com/skxeve/PersonalLineBot/line/log"
	"google.golang.org/appengine"
	"net/http"
	"os"
)

type HttpContext struct {
	Logger *l.Logger
}

func GetHttpContext(r *http.Request) *HttpContext {
	gae_instance := os.Getenv("GAE_INSTANCE")
	c := new(HttpContext)
	if gae_instance != "" {
		c.Logger = &l.Logger{
			Env:     1,
			Context: appengine.NewContext(r),
		}
	} else {
		c.Logger = &l.Logger{}
	}
	return c
}
