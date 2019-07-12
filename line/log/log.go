package log

import (
	"golang.org/x/net/context"
	gaelog "google.golang.org/appengine/log"
	"log"
)

type Logger struct {
	Env     int
	Context context.Context
}

const LOGGER_ENV_GAE = 1

func (l Logger) Debugf(format string, v ...interface{}) {
	switch l.Env {
	case LOGGER_ENV_GAE:
		gaelog.Debugf(l.Context, format, v...)
	default:
		log.Printf("[DEBUG] "+format, v...)
	}
}

func (l Logger) Infof(format string, v ...interface{}) {
	switch l.Env {
	case LOGGER_ENV_GAE:
		gaelog.Infof(l.Context, format, v...)
	default:
		log.Printf("[INFO] "+format, v...)
	}
}

func (l Logger) Warningf(format string, v ...interface{}) {
	switch l.Env {
	case LOGGER_ENV_GAE:
		gaelog.Warningf(l.Context, format, v...)
	default:
		log.Printf("[WARNING] "+format, v...)
	}
}

func (l Logger) Errorf(format string, v ...interface{}) {
	switch l.Env {
	case LOGGER_ENV_GAE:
		gaelog.Errorf(l.Context, format, v...)
	default:
		log.Printf("[ERROR] "+format, v...)
	}
}
