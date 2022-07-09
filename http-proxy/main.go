package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

const (
	PORT            = "2770"
	DESTINATION_URI = "http://localhost:2200/dump"
)

func getDestination() string {
	return DESTINATION_URI
}

func main() {
	fmt.Printf("Starting http proxy on port: %s\n", PORT)
	http.HandleFunc("/", handleProxyRequest)

	log.Fatal(http.ListenAndServe(":"+PORT, nil))
}

func handleProxyRequest(writer http.ResponseWriter, request *http.Request) {
	path := request.URL.Path
	proxyURL, err := url.Parse(DESTINATION_URI + path)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Proxy to URI: %s  Proxy path: %s\n", proxyURL, path)

	proxy := httputil.NewSingleHostReverseProxy(proxyURL)
	//writer.Header().Add("test", "value")
	proxy.ServeHTTP(writer, request)
}
