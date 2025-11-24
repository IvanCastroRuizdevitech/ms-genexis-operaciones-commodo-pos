package infrastructure_external_nethttp_nethttp

import (
	"bytes"
	"io"
	"net/http"
	"time"
)

type NetHTTPClient struct{}

func (c *NetHTTPClient) Get(url string, headers map[string]string) ([]byte, error) {
	_, body, err := c.Request(http.MethodGet, url, nil, headers)
	return body, err
}

func (c *NetHTTPClient) Post(url string, bodyRequest []byte, headers map[string]string) ([]byte, error) {
	_, body, err := c.Request(http.MethodPost, url, bodyRequest, headers)
	return body, err
}

func (c *NetHTTPClient) Put(url string, bodyRequest []byte, headers map[string]string) ([]byte, error) {
	_, body, err := c.Request(http.MethodPut, url, bodyRequest, headers)
	return body, err
}

func (c *NetHTTPClient) Request(method string, url string, body []byte, headers map[string]string) (int, []byte, error) {
	var reader io.Reader
	if len(body) > 0 {
		reader = bytes.NewBuffer(body)
	}

	req, err := http.NewRequest(method, url, reader)
	if err != nil {
		return 0, nil, err
	}

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	if req.Header.Get("Content-Type") == "" && (method == http.MethodPost || method == http.MethodPut || method == http.MethodPatch) {
		req.Header.Set("Content-Type", "application/json")
	}

	client := &http.Client{
		Timeout: 60 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()

	bodyResponse, err := io.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, nil, err
	}

	return resp.StatusCode, bodyResponse, nil
}
