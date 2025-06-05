package httpclient

import (
	"bytes"
	"io"
	"net/http"
	"time"
)

var client = &http.Client{Timeout: 5 * time.Second}

func Get(url string) ([]byte, error) {
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

func Post(url string, contentType string, body []byte) ([]byte, error) {
	resp, err := client.Post(url, contentType, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}
