package httpserver

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type server struct {
	extSrvURL *url.URL
}

// serveSimpleHTTP is an handler for an dummy http request
func serveSimpleHTTP(w http.ResponseWriter, r *http.Request) {
	if _, ok := r.URL.Query()["fail"]; ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Fprint(w, "Hello world!")
	w.WriteHeader(http.StatusOK)
}

// serveHTTP is an handler for an http request
func (s *server) ServeHTTPToExternal(w http.ResponseWriter, r *http.Request) {
	if _, ok := r.URL.Query()["fail"]; ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resp, err := http.Get(s.extSrvURL.String())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if resp.StatusCode >= 400 {
		http.Error(w, resp.Status, resp.StatusCode)
		return
	}
	b, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// do something with b
	// here we just return it as it
	w.Write(b)
	w.WriteHeader(http.StatusOK)
}
