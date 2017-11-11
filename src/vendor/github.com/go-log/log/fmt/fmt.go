package fmt

import (
	"fmt"

	"github.com/go-log/log"
)

type fmtLogger struct{}

var (
	_ log.Logger = New()
)

func (t *fmtLogger) Log(v ...interface{}) {
	fmt.Print(v...)
}

func (t *fmtLogger) Logf(format string, v ...interface{}) {
	fmt.Printf(format, v...)
}

func New() *fmtLogger {
	return &fmtLogger{}
}
