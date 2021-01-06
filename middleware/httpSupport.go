package middleware

import (
	"bufio"
	"github.com/eavesmy/puppy"
	"net/http"
)

func HttpSupport(ctx *puppy.Context) error {

	req, err := http.ReadRequest(ctx.Body.(*bufio.Reader))
	if err != nil {
		return nil
	}

	ctx.Req = req
	ctx.Res = puppy.Res{
		Headers:    http.Header{},
		Context:    ctx,
		Status:     "200 OK",
		StatusCode: 200,
		Proto:      "HTTP/1.0",
		ProtoMajor: 1,
		ProtoMinor: 0,
	}
	ctx.IsHttp = true

	return nil
}
