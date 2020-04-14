package httputil

import (
	"bytes"
	"encoding/json"
	"github.com/google/go-querystring/query"
	"io/ioutil"
	"net/http"
)

func Get(paras GetParasTuple) (data []byte, err error) {
	httpClient := getHttpClient(paras.Client)
	req, err := http.NewRequest(http.MethodGet, paras.URL, nil)
	if err != nil {
		return
	}
	for k, v := range paras.Headers {
		req.Header.Set(k, v)
	}
	if paras.Host != "" {
		req.Host = paras.Host
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		err = StatusError{resp.StatusCode}
		return
	}
	data, err = ioutil.ReadAll(resp.Body)
	return
}

func GetForm(paras GetFormParasTuple) (data []byte, err error) {
	httpClient := getHttpClient(paras.Client)
	req, err := http.NewRequest(http.MethodGet, paras.URL, nil)
	if err != nil {
		return
	}

	v, err := query.Values(paras.Req)
	if err != nil {
		return
	} else {
		req.URL.RawQuery = v.Encode()
	}

	for k, v := range paras.Headers {
		req.Header.Set(k, v)
	}
	if req.Header.Get(ContentType) == "" {
		req.Header.Set(ContentType, AppForm)
	}
	if paras.Host != "" {
		req.Host = paras.Host
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		err = StatusError{resp.StatusCode}
		return
	}
	data, err = ioutil.ReadAll(resp.Body)
	return
}

func Post(paras PostParasTuple) (data []byte, err error) {
	buffer := new(bytes.Buffer)
	if err = json.NewEncoder(buffer).Encode(paras.Req); err != nil {
		return
	}
	httpClient := getHttpClient(paras.Client)
	req, err := http.NewRequest(http.MethodPost, paras.URL, buffer)
	if err != nil {
		return
	}
	for k, v := range paras.Headers {
		req.Header.Set(k, v)
	}
	if paras.Host != "" {
		req.Host = paras.Host
	}
	req.Header.Set(ContentType, AppJson)
	resp, err := httpClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusCreated {
		data, err = ioutil.ReadAll(resp.Body)
		return
	}

	err = StatusError{resp.StatusCode}
	return
}

func PostForm(paras PostFormParasTuple) (data []byte, err error) {
	v, err := query.Values(paras.Req)
	if err != nil {
		return
	}

	buffer := bytes.NewBufferString(v.Encode())
	httpClient := getHttpClient(paras.Client)
	req, err := http.NewRequest(http.MethodPost, paras.URL, buffer)
	if err != nil {
		return
	}
	for k, v := range paras.Headers {
		req.Header.Set(k, v)
	}

	if req.Header.Get(ContentType) == "" {
		req.Header.Set(ContentType, AppForm)
	}

	if paras.Host != "" {
		req.Host = paras.Host
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		err = StatusError{resp.StatusCode}
		return
	}
	data, err = ioutil.ReadAll(resp.Body)
	return
}

func PostBytes(paras PostBytesTuple) (data []byte, err error) {
	httpClient := getHttpClient(paras.Client)
	req, err := http.NewRequest(http.MethodPost, paras.URL, bytes.NewBuffer(paras.Req))
	if err != nil {
		return
	}

	for k, v := range paras.Headers {
		req.Header.Set(k, v)
	}
	if paras.Host != "" {
		req.Host = paras.Host
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = StatusError{resp.StatusCode}
		return
	}
	data, err = ioutil.ReadAll(resp.Body)
	return
}

func Delete(paras GetParasTuple) (data []byte, err error) {
	httpClient := getHttpClient(paras.Client)
	req, err := http.NewRequest(http.MethodDelete, paras.URL, nil)
	if err != nil {
		return
	}
	for k, v := range paras.Headers {
		req.Header.Set(k, v)
	}
	if paras.Host != "" {
		req.Host = paras.Host
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		data, err = ioutil.ReadAll(resp.Body)
		return
	}
	err = StatusError{resp.StatusCode}
	return
}
