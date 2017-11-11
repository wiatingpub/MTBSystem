package appengine

import (
	"testing"

	"github.com/go-log/log"
	"google.golang.org/appengine/aetest"
)

func testLog(l log.Logger) {
	l.Log("test\n")
}

func testLogf(l log.Logger) {
	l.Logf("%s", "test\n")
}

func TestLogLogger(t *testing.T) {
	l := new(appengineLogger)
	ctx, done, err := aetest.NewContext()
	if err != nil {
		t.Fatal(err)
	}
	defer done()
	l.context = ctx
	testLog(l)
	testLogf(l)
}
