package puppy

import (
	"fmt"
	"net"
	"net/http"
)

type Server struct {
	conn net.Conn
	http bool
	core *Puppy
}

func (s *Server) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	ctx := NewHttpContext(req, res)
	s.core.mds.run(ctx)
}

func (s *Server) Listen(p *Puppy) (err error) {

	s.core = p

	isHttp := p.Conf.Http

	ln, err := net.Listen("tcp", p.Conf.addr)

	if err != nil {
		panic(err)
	}

	fmt.Println("server start at", p.Conf.addr)

	if isHttp {
		ln.Addr()
		err = http.Serve(ln, s)
	} else {
		s.Serve(ln)
	}

	if err != nil {
		return
	}

	return
	// http 使用系统提供的 http 方法
}

func (s *Server) Serve(ln net.Listener) {

	for {
		conn, err := ln.Accept()

		if err != nil {
			// 请求失败
			continue
		}

		ctx := NewContext(conn, s)
		s.core.mds.run(ctx)
	}
}
