package middleware

import (
	"bufio"
	"github.com/eavesmy/puppy"
	"github.com/go-http-utils/cookie"
	"net/http"
)

func HttpSupport(ctx *puppy.Context) error {

	req, err := http.ReadRequest(ctx.Body.(*bufio.Reader))
	if err != nil {
		return nil
	}

	ctx.Req = req
	ctx.Method = req.Method
	ctx.Path = req.RequestURI

	ctx.Res = puppy.Res{
		Headers:    http.Header{},
		Context:    ctx,
		Status:     "200 OK",
		StatusCode: 200,
		Proto:      "HTTP/1.0",
		ProtoMajor: 1,
		ProtoMinor: 0,
	}

	ctx.Cookies = cookie.New(ctx.Res, ctx.Req)

	ctx.IsHttp = true

	return nil
}
