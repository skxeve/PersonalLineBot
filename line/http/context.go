package http

import (
	l "github.com/skxeve/PersonalLineBot/line/log"
	"net/http"
	"os"
)

type HttpContext struct {
	Logger  *l.Logger
	Request *http.Request
}

func NewContext(r *http.Request) *HttpContext {
	gae_instance := os.Getenv("GAE_INSTANCE")
	c := new(HttpContext)
	c.Request = r
	if gae_instance != "" {
		c.Logger = &l.Logger{
			Env:     1,
			Primary: os.Getenv("GAE_INSTANCE")[:7],
		}
	} else {
		c.Logger = &l.Logger{
			Env: 0,
		}
	}
	return c
}
