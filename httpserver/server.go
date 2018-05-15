package httpserver

import (
	"net/http"
	"net/url"
)

func runServer(extServiceURL *url.URL) error {
	s := &server{extSrvURL: extServiceURL}
	http.HandleFunc("/ext", s.ServeHTTPToExternal)
	http.HandleFunc("/simple", s.ServeSimpleHTTP)
	return http.ListenAndServe(":12345", nil)
}
