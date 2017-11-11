package fmt

import (
	"testing"

	"github.com/go-log/log"
)

func testLog(l log.Logger) {
	l.Log("test\n")
}

func testLogf(l log.Logger) {
	l.Logf("%s", "test\n")
}

func TestFMTLogger(t *testing.T) {
	l := new(fmtLogger)
	testLog(l)
	testLogf(l)
}
