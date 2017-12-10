package apid

import (
	"github.com/micro/go-micro/metadata"
	"golang.org/x/net/context"
	"net/http"
	"path"
	"regexp"
	"strings"
)

var (
	versionRe = regexp.MustCompilePOSIX("^v[0-9]+$")
)

// Translates /foo/bar/zool into api service go.micro.api.foo method Bar.Zool
// Translates /foo/bar into api service go.micro.api.foo method Foo.Bar
func PathToReceiver(ns, p string) (string, string) {
	p = path.Clean(p)
	p = strings.TrimPrefix(p, "/")
	parts := strings.Split(p, "/")

	// If we've got two or less parts
	// Use first part as service
	// Use all parts as method
	if len(parts) <= 2 {
		service := ns + strings.Join(parts[:len(parts)-1], ".")
		method := strings.Title(strings.Join(parts, "."))
		return service, method
	}

	// Treat /v[0-9]+ as versioning where we have 3 parts
	// /v1/foo/bar => service: v1.foo method: Foo.bar
	if len(parts) == 3 && versionRe.Match([]byte(parts[0])) {
		service := ns + strings.Join(parts[:len(parts)-1], ".")
		method := strings.Title(strings.Join(parts[len(parts)-2:], "."))
		return service, method
	}

	// Service is everything minus last two parts
	// Method is the last two parts
	service := ns + strings.Join(parts[:len(parts)-2], ".")
	method := strings.Title(strings.Join(parts[len(parts)-2:], "."))
	return service, method
}

func RequestToContext(r *http.Request) context.Context {
	ctx := context.Background()
	md := make(metadata.Metadata)
	for k, v := range r.Header {
		md[k] = strings.Join(v, ",")
	}
	return metadata.NewContext(ctx, md)
}
