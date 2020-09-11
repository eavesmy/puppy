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

	Body       io.Reader
	RemoteAddr string
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

func (c *Context) Call()   {}
func (c *Context) ReCall() {}
func (c *Context) Json()   {}
func (c *Context) Text()   {}
