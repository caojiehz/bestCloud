package httputil

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

type TFunc func(w http.ResponseWriter, r *http.Request)

func StartTestServer(t TFunc) {
	mux := http.NewServeMux()
	mux.HandleFunc("/test", t)
	s := &http.Server{
		Addr:           "127.0.0.1:9527",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 10 << 20,
	}

	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}
}

func TestHttpClient404(t *testing.T) {
	go StartTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
	})
	time.Sleep(1 * time.Second)
	_, err := RetryPostBytes(PostBytesTuple{
		URL:      "http://127.0.0.1:9527/test",
		Req:      []byte("hello"),
		RetryNum: 3,
	})

	assert.NotEqual(t, err, nil)
}
