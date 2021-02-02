package puppy

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/eavesmy/golang-lib/context"
	"github.com/go-http-utils/cookie"
	"io"
	"net/url"
	// "io/ioutil"
	"net"
	"net/http"
)

type Context struct {
	conn net.Conn

	ctx    *context.Ctx
	server *Server

	Req     *http.Request
	Res     http.ResponseWriter
	Cookies *cookie.Cookies
	Query   url.Values

	Method string
	Path   string
	Length int64

	Body       io.Reader
	RemoteAddr string

	IsHttp bool

	buffer_read []byte
	read        int
	buffer_pool []byte

	ParseBodyByHand func(i interface{}) error
	EventClose      func(*Context) error
	EventAfterClose func(*Context) error
}

func NewContext(conn net.Conn, server *Server) *Context {
	return &Context{
		conn:        conn,
		server:      server,
		Body:        bufio.NewReader(conn),
		ctx:         context.New(),
		buffer_read: make([]byte, server.core.Conf.BufferSize),
		buffer_pool: []byte{},
		read:        0,
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

// Default parse method. Just parse json format.
func (c *Context) ParseBody(i interface{}) error {

	if c.ParseBodyByHand != nil {
		return c.ParseBodyByHand(i)
	}

	index, err := c.Body.Read(c.buffer_read[c.read:])
	c.read += index

	if err != nil {
		// error
		return err
	}

	if len(c.buffer_pool) > 0 {
		c.buffer_pool = append(c.buffer_pool, c.buffer_read[c.read:]...)

		err := read(c.buffer_pool, &i)

		if err != nil { // try read
			c.buffer_pool = []byte{}
			c.buffer_read = make([]byte, c.server.core.Conf.BufferSize)
			c.read = 0

			return nil
		}
	}

	err = json.Unmarshal(c.buffer_read[:c.read], &i)

	if err != nil {
		// Wait next data.
		fmt.Println(err)
	} else {
		c.buffer_read = make([]byte, c.server.core.Conf.BufferSize)
		c.read = 0
	}

	if c.read >= c.server.core.Conf.BufferSize {
		c.read = 0
		c.buffer_pool = append(c.buffer_pool, c.buffer_read[c.read:]...)
		c.buffer_read = make([]byte, c.server.core.Conf.BufferSize)
	}

	return nil
}

func read(b []byte, i interface{}) error {
	return json.Unmarshal(b, i)
}

func (c *Context) Call(pattern string, arg interface{}) {
	// 查找对应路由，拿到节点
	// 校验参数,判断参数
}

func (c *Context) ReCall() {}

// 关闭当前链接
func (c *Context) Close() {

	if c.EventClose != nil {
		c.EventClose(c)
	}

	c.conn.Close() // 关闭 conn
	c.ctx.Cancel()

	if c.EventAfterClose != nil {
		c.EventAfterClose(c)
	}
}

func (c *Context) Json(i interface{}, statusCodes ...int) (err error) {
	statusCode := 200
	if len(statusCodes) > 0 {
		statusCode = statusCodes[0]
	}

	b, err := json.Marshal(i)
	if err != nil {
		return
	}

	c.Res.Header().Add("Content-Type", "application/json")

	if c.Res != nil {
		c.Res.WriteHeader(statusCode)
		_, err = c.Res.Write(b)
	} else if c.conn != nil {
		_, err = c.conn.Write(b)
	}

	return
}

func (c *Context) Text(text string, statusCodes ...int) (err error) {

	statusCode := 200
	if len(statusCodes) > 0 {
		statusCode = statusCodes[0]
	}

	if c.Res != nil {
		c.Res.WriteHeader(statusCode)
		_, err = c.Res.Write([]byte(text))
	} else if c.conn != nil {
		_, err = c.conn.Write([]byte(text))
	}
	return
}

// Support Res.Write
// Just Http.
func (c *Context) Write(d []byte) (i int, err error) {

	statusCode := 200
	if s_statusCode := c.Get("StatusCode"); s_statusCode != nil {
		statusCode = s_statusCode.(int)
	}

	statusText := http.StatusText(statusCode)
	status := fmt.Sprintf("%d", statusCode) + " " + statusText

	// Header
	if c.Res.Header().Get("Content-Type") == "" {
		c.Res.Header().Add("Content-Type", http.DetectContentType(d))
	}
	c.Res.Header().Add("server", "puppy/1.0.0")

	buffer := bytes.NewBuffer([]byte{})
	c.Res.Header().Write(buffer)

	// Cookies

	res := []byte{}
	res = append(res, []byte(c.Res.(Res).Proto+" "+status+"\n")...)
	res = append(res, buffer.Bytes()...)
	res = append(res, []byte("\n\r")...)
	res = append(res, d...)

	i, err = c.conn.Write(res)
	if c.IsHttp {
		c.conn.Close()
	}

	return
}
