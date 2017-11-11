package logrus

import (
	"github.com/go-log/log"
	"github.com/sirupsen/logrus"
)

type logrusLogger struct {
	*logrus.Entry
}

var (
	_ log.Logger = New()
)

func (l *logrusLogger) Log(v ...interface{}) {
	if l.Entry != nil {
		l.Entry.Print(v...)
	} else {
		logrus.Print(v...)
	}
}

func (l *logrusLogger) Logf(format string, v ...interface{}) {
	if l.Entry != nil {
		l.Entry.Printf(format, v...)
	} else {
		logrus.Printf(format, v...)
	}
}

func WithFields(f logrus.Fields) log.Logger {
	return &logrusLogger{logrus.WithFields(f)}
}

func New() *logrusLogger {
	return &logrusLogger{}
}
