package util

import (
	"io"
	"net/http"
)

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type HttpClientMock struct {
	StatusCode int
	Body       io.ReadCloser
	Err        error
}

func (h *HttpClientMock) Do(req *http.Request) (*http.Response, error) {
	return &http.Response{
		StatusCode: h.StatusCode,
		Body:       h.Body,
	}, h.Err
}
