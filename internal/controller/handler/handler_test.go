package handler

import (
	"errors"
	"net/http"
)

type dummyResponseWriter struct {
	Code int
}

func (d *dummyResponseWriter) Header() http.Header {
	return http.Header{}
}

func (d *dummyResponseWriter) Write([]byte) (int, error) {
	return 0, errors.New("write failed")
}

func (d *dummyResponseWriter) WriteHeader(statusCode int) {
	d.Code = statusCode
}

type bodyReader struct {
}

func (b *bodyReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("read failed")
}
