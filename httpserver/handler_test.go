package httpserver

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_server_ServeSimpleHTTP(t *testing.T) {
	req := httptest.NewRequest("GET", "http://sample.io/simple", nil)
	w := httptest.NewRecorder()
	serveSimpleHTTP(w, req)
	resp := w.Result()
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	b, err := ioutil.ReadAll(resp.Body)
	require.NoError(t, err)
	assert.Equal(t, "Hello world!", string(b))

	req = httptest.NewRequest("GET", "http://sample.io/simple?fail", nil)
	w = httptest.NewRecorder()
	serveSimpleHTTP(w, req)
	resp = w.Result()
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	b, err = ioutil.ReadAll(resp.Body)
	require.NoError(t, err)
	assert.Empty(t, string(b))
}

func Test_server_ServeHTTPToExternal(t *testing.T) {

	tests := []struct {
		name            string
		handler         http.Handler
		urlExtras       string
		expectedStatus  int
		expectedContent string
	}{
		{"CheckOK", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, "all good")
		}), "", http.StatusOK, "all good"},
		{"CheckFail", nil, "?fail", http.StatusBadRequest, ""},
		{"CheckExtError", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		}), "", http.StatusUnauthorized, "401 Unauthorized\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			ts := httptest.NewServer(tt.handler)
			defer ts.Close()
			u, err := url.Parse(ts.URL)
			require.NoError(t, err)
			s := &server{
				extSrvURL: u,
			}
			req := httptest.NewRequest("GET", fmt.Sprintf("http://sample.io/ext%s", tt.urlExtras), nil)
			w := httptest.NewRecorder()

			s.ServeHTTPToExternal(w, req)

			resp := w.Result()
			b, err := ioutil.ReadAll(resp.Body)
			require.NoError(t, err)

			assert.Equal(t, tt.expectedStatus, resp.StatusCode)
			assert.Equal(t, tt.expectedContent, string(b))
		})
	}
}
