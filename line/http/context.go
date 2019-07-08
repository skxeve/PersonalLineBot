package http

import (
	l "../log"
)

type HttpContext struct {
	Logger *l.Logger
}

func GetHttpContext() *HttpContext {
	c := new(HttpContext)
	c.Logger = &l.Logger{}
	return c
}
