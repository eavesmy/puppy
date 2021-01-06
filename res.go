package puppy

import (
	"net/http"
)

type Res struct {
	Headers http.Header
	Context *Context

	Status     string // e.g. "200 OK"
	StatusCode int    // e.g. 200
	Proto      string // e.g. "HTTP/1.0"
	ProtoMajor int    // e.g. 1
	ProtoMinor int    // e.g. 0
}

func (r Res) Write(d []byte) (int, error) {
	return r.Context.Write(d)
}

func (r Res) Header() http.Header {
	return r.Headers
}

func (r Res) SetHeader(k, v string) {
	r.Headers.Add(k, v)
}

func (r Res) WriteHeader(statusCode int) {
	r.Context.Set("StatusCode", statusCode)
}
