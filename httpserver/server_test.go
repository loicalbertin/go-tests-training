package httpserver

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_server_runServer(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hi there!\n")
	}))
	defer ts.Close()
	u, err := url.Parse(ts.URL)
	require.NoError(t, err)

	err = runServer(context.Background(), u)
	assert.Error(t, err, "expecting an error when running http server with a not cancellable context")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	err = runServer(ctx, u)
	require.NoError(t, err)

	time.Sleep(30 * time.Millisecond)
	resp, err := http.DefaultClient.Get("http://127.0.0.1:12345/simple")
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	assert.Equal(t, "Hello world!", string(b))

	resp, err = http.DefaultClient.Get("http://127.0.0.1:12345/ext")
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	b, err = ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	assert.Equal(t, "Hi there!\n", string(b))

}
