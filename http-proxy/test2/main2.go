package test2

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

type ProxyData struct {
	Writer  http.ResponseWriter
	Request *http.Request
}

func main() {
	fmt.Printf("Starting http proxy on port: %s\n", PORT)

	trigger("/", PORT)
}

func trigger(path string, port string) {
	http.HandleFunc(path, handleProxyRequest)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func activity(pd ProxyData) {
	path := pd.Request.URL.Path
	proxyURL, err := url.Parse(DESTINATION_URI + path)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Proxy to URI: %s  Proxy path: %s\n", proxyURL, path)

	proxy := httputil.NewSingleHostReverseProxy(proxyURL)
	pd.Writer.Header().Add("test", "value") //response headers
	proxy.ServeHTTP(pd.Writer, pd.Request)
}

func handleProxyRequest(writer http.ResponseWriter, request *http.Request) {

}
