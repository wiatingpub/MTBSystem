package appengine

import (
	"errors"

	"golang.org/x/net/context"
	appenginelog "google.golang.org/appengine/log"
)

type appengineLogger struct {
	context context.Context
}

func (a *appengineLogger) Log(v ...interface{}) {
	appenginelog.Debugf(a.context, "%v", v...)
}

func (a *appengineLogger) Logf(format string, v ...interface{}) {
	appenginelog.Debugf(a.context, format, v...)
}

func New(ctx context.Context) (*appengineLogger, error) {
	if ctx == nil {
		return nil, errors.New("appengine context required")
	}
	return &appengineLogger{context: ctx}, nil
}
