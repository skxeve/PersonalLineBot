package log

import (
	"golang.org/x/net/context"
	alog "google.golang.org/appengine/log"
	"log"
)

type Logger struct {
	Env     int
	Context interface{}
}

const LOG_ENV_GAE = 1

func (l Logger) Debugf(format string, v ...interface{}) {
	switch l.Env {
	case LOG_ENV_GAE:
		alog.Debugf(l.Context.(context.Context), format, v...)
	default:
		log.Printf("[DEBUG] "+format, v...)
	}
}

func (l Logger) Infof(format string, v ...interface{}) {
	switch l.Env {
	case LOG_ENV_GAE:
		alog.Infof(l.Context.(context.Context), format, v...)
	default:
		log.Printf("[INFO] "+format, v...)
	}
}

func (l Logger) Warningf(format string, v ...interface{}) {
	switch l.Env {
	case LOG_ENV_GAE:
		alog.Warningf(l.Context.(context.Context), format, v...)
	default:
		log.Printf("[WARNING] "+format, v...)
	}
}

func (l Logger) Errorf(format string, v ...interface{}) {
	switch l.Env {
	case LOG_ENV_GAE:
		alog.Errorf(l.Context.(context.Context), format, v...)
	default:
		log.Printf("[ERROR] "+format, v...)
	}
}
