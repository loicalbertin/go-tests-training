package httpserver

import (
	"context"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

func runServer(ctx context.Context, extServiceURL *url.URL) error {
	ch := ctx.Done()
	if ch == nil {
		return errors.Errorf("A cancelable context is required")
	}

	s := &server{extSrvURL: extServiceURL}
	http.HandleFunc("/ext", s.ServeHTTPToExternal)
	http.HandleFunc("/simple", serveSimpleHTTP)

	hs := &http.Server{Addr: ":12345"}
	go hs.ListenAndServe()
	go func() {
		<-ctx.Done()
		hs.Close()
	}()
	return nil
}
