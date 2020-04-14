package httputil

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"time"
)

const (
	ContentType = "Content-Type"
	AppJson     = "application/json;charset=UTF-8"
	AppForm     = "application/x-www-form-urlencoded;charset=UTF-8"
)

var client *http.Client

func init() {
	client = &http.Client{
		Timeout: 3 * time.Second, //client强求超时
		Transport: &http.Transport{
			MaxIdleConnsPerHost: 1000,
			TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
		},
	}
}

func getHttpClient(c *http.Client) *http.Client {
	if c != nil {
		return c
	} else {
		return client
	}
}

type GetParasTuple struct {
	Client   *http.Client
	URL      string
	Host     string
	RetryNum int
	Headers  map[string]string
	Resp     interface{}
}

type GetFormParasTuple struct {
	Client   *http.Client
	URL      string
	Host     string
	RetryNum int
	Headers  map[string]string
	Req      interface{}
	Resp     interface{}
}

type PostParasTuple struct {
	Client   *http.Client
	URL      string
	Host     string
	Req      interface{}
	Resp     interface{}
	RetryNum int
	Headers  map[string]string
}

type PostFormParasTuple struct {
	Client   *http.Client
	URL      string
	Host     string
	Req      interface{}
	Resp     interface{}
	RetryNum int
	Headers  map[string]string
}

type PostBytesTuple struct {
	Client   *http.Client
	URL      string
	Host     string
	Req      []byte
	RetryNum int
	Headers  map[string]string
}

func GetHeaders(req *http.Request, keys []string) map[string]string {
	ret := make(map[string]string)
	for _, key := range keys {
		value := req.Header.Get(key)
		if value == "" {
			continue
		}
		ret[key] = value
	}
	return ret
}

type StatusError struct {
	Code int
}

func (s StatusError) Error() string {
	return fmt.Sprintf("http status code %d", s.Code)
}
