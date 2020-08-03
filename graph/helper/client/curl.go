package client

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type CURLOption func(*CURLOpts)

type CURLOpts struct {
	Headers map[string]string
	Body    interface{}
}

func (cs S) CURL(method string, url string, opts ...CURLOption) (interface{}, error) {
	var result interface{}
	curlOpt := &CURLOpts{}
	for _, opt := range opts {
		opt(curlOpt)
	}

	byteBody, _ := json.Marshal(curlOpt.Body)
	req, errReq := http.NewRequest(method, url, bytes.NewBuffer(byteBody))
	if errReq != nil {
		return nil, errReq
	}

	for keyHeader, valueHeader := range curlOpt.Headers {
		req.Header.Set(keyHeader, valueHeader)
	}

	client := &http.Client{}
	resp, errResp := client.Do(req)
	if errResp != nil {
		return nil, errResp
	}
	defer func() { _ = resp.Body.Close() }()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(respBody, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (cs S) WithBody(body interface{}) CURLOption {
	return func(curlOpts *CURLOpts) {
		curlOpts.Body = body
	}
}

func (cs S) WithHeader(header map[string]string) CURLOption {
	return func(curlOpts *CURLOpts) {
		curlOpts.Headers = header
	}
}
