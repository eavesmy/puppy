package puppy

import (
	"bufio"
	"github.com/eavesmy/golang-lib/context"
	"io"
	"net"
	"net/http"
)

type Context struct {
	conn net.Conn

	ctx    *context.Ctx
	server *Server

	Req *http.Request
	Res http.ResponseWriter

	Method string
	Path   string
	Length int64

	Body       io.Reader
	RemoteAddr string
}

func NewHttpContext(req *http.Request, res http.ResponseWriter) *Context {
	return &Context{
		Body:       req.Body,
		ctx:        context.New(),
		Req:        req,
		Res:        res,
		Method:     req.Method,
		RemoteAddr: req.RemoteAddr,
		Path:       req.RequestURI,
		Length:     req.ContentLength,
	}
}

func NewContext(conn net.Conn, server *Server) *Context {
	return &Context{
		conn:   conn,
		server: server,
		Body:   bufio.NewReader(conn),
		ctx:    context.New(),
	}
}

func (c *Context) SetHeader(key, value string) {
	if c.Res != nil {
		c.Res.Header().Set(key, value)
	}
}

func (c *Context) GetHeader(key string) string {
	if c.Req != nil {
		return c.Req.Header.Get(key)
	}
	return ""
}

func (c *Context) Set(k, v interface{}) {
	c.ctx.Set(k, v)
}

func (c *Context) Get(k interface{}) interface{} {
	return c.ctx.Get(k)
}

func (c *Context) ParseBody() {

}

func (c *Context) Call(pattern string, arg interface{}) {
	// 查找对应路由，拿到节点
	// 校验参数,判断参数
}

func (c *Context) ReCall() {}

func (c *Context) Json() {}

func (c *Context) Text(text string, statusCodes ...int) (err error) {

	statusCode := 200
	if len(statusCodes) > 0 {
		statusCode = statusCodes[2]
	}

	if c.Res != nil {
		c.Res.WriteHeader(statusCode)
		_, err = c.Res.Write([]byte(text))
	} else if c.conn != nil {
		_, err = c.conn.Write([]byte(text))
	}
	return
}
