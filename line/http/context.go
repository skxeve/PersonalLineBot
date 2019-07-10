package http

import (
	l "github.com/skxeve/PersonalLineBot/line/log"
)

type HttpContext struct {
	Logger *l.Logger
}

func GetHttpContext() *HttpContext {
	c := new(HttpContext)
	c.Logger = &l.Logger{}
	return c
}
