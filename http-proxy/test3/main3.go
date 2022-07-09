package test3

import (
	"fmt"
	"github.com/elazarl/goproxy"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func start() {
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true
	log.Fatal(http.ListenAndServe(":8080", proxy))

	proxy.OnRequest().DoFunc(
		func(req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
			fmt.Printf("Req URL: %s", req.URL.Path)
			return req, nil
		})

	proxy.OnRequest().DoFunc(
		func(req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
			url, _ := url.Parse("https://httpbin.org/status/200")
			proxy := httputil.NewSingleHostReverseProxy(url)
			//writer.Header().Add("test", "value")
			proxy.ServeHTTP(nil, req)
			return req, nil
		})

}
