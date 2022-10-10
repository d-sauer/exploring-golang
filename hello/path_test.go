package main

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"net/url"
	"os"
	"path"
	"strings"
	"testing"
)

func TestPath(t *testing.T) {
	var s = "/hello-oas/openapi-spec/ServiceHelloV1/specification"
	dir := path.Dir(s)
	path := path.Join(dir, path.Base(s))

	fmt.Printf("path: %s", path)
}

func TestPath2(t *testing.T) {
	var s = "/hello-oas/openapi-spec/ServiceHelloV1/specification/openapi.yaml"
	var r = ""

	var p = s
	if r != "" {
		p = path.Join(path.Dir(s), r)
	}

	fmt.Printf("path: %s", p)
}

func TestPath3(t *testing.T) {
	var s = "/ServiceHelloV1/specification"
	var f = path.Join(s, string(os.PathSeparator))

	fmt.Printf("path: %s", f)
}

func TestJoinURL(t *testing.T) {
	var a = &url.URL{Path: "/hello-oas/openapi-spec/ServiceHelloV1/specification"}
	var b = &url.URL{Path: "./openapi.yaml"}
	var r, err = url.JoinPath(a.Path, b.Path)
	fmt.Printf("path: %s", r)

	a = &url.URL{Path: "hello-oas/openapi-spec/ServiceHelloV1/specification/foo.yaml"}
	b = &url.URL{Path: "./openapi.yaml"}
	var n = a.ResolveReference(b)
	fmt.Printf("\npath: %s", n)

	a = &url.URL{Path: "hello-oas/openapi-spec/ServiceHelloV1/specification/foo.yaml"}
	b = &url.URL{Path: "./openapi.yaml"}
	fmt.Printf("\npath: %s, %s", a.EscapedPath(), b.EscapedPath())

	require.NoError(t, err)
}

func TestIsResource(t *testing.T) {
	var a = isResourcePath(&url.URL{Path: "hello-oas/openapi-spec/ServiceHelloV1/specification/foo.yaml"})
	fmt.Printf("\nisResource: %t", a)

	a = isResourcePath(&url.URL{Path: "hello-oas/openapi-spec/ServiceHelloV1/specification"})
	fmt.Printf("\nisResource: %t", a)

	a = isResourcePath(&url.URL{Path: "hello-oas/openapi-spec/ServiceHelloV1/specification/"})
	fmt.Printf("\nisResource: %t", a)
}

func isResourcePath(base *url.URL) bool {
	// check if there is file extension after last slash
	var s = strings.LastIndex(base.Path, "/")
	var n = strings.LastIndex(base.Path, ".")
	return n > s
}
