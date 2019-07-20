package log

import (
	"fmt"
	"golang.org/x/net/context"
	"log"
	"os"
)

type Logger struct {
	Env     int
	Context context.Context
}

const LOGGER_ENV_GAE = 1

func (l Logger) Debugf(format string, v ...interface{}) {
	l.logf(format, "DEBUG", v...)
}

func (l Logger) Infof(format string, v ...interface{}) {
	l.logf(format, "INFO", v...)
}

func (l Logger) Warningf(format string, v ...interface{}) {
	l.logf(format, "WARNING", v...)
}

func (l Logger) Errorf(format string, v ...interface{}) {
	l.logf(format, "ERROR", v...)
}

func (l Logger) logf(format, level string, v ...interface{}) {
	switch l.Env {
	case LOGGER_ENV_GAE:
		l.gaeLogf(format, level, v...)
	default:
		l.defaultLogf(format, level, v...)
	}
}

func (l Logger) defaultLogf(format, level string, v ...interface{}) {
	log.Printf(
		"[%s] %s",
		level,
		fmt.Sprintf(format, v...))
}

func (l Logger) gaeLogf(format, level string, v ...interface{}) {
	log.Printf(
		"[%s:%s] %s <%s>",
		level,
		os.Getenv("GAE_INSTANCE"),
		fmt.Sprintf(format, v...),
		os.Getenv("GAE_VERSION"))
}
