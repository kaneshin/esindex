package elasticsearch

import (
	"io"
	"net/http"
)

type Request struct {
	*http.Request
	method string
	path   string
	body   io.Reader
}

func NewRequest(method, path string, body io.Reader) *Request {
	return &Request{
		method: method,
		path:   path,
		body:   body,
	}
}
