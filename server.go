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

}

func (s *Server) Listen(p *Puppy) (err error) {

	isHttp := p.Conf.Http

	protocol := ""

	if isHttp {
		protocol = "http"
	} else {
		protocol = "tcp"
	}

	ln, err := net.Listen(protocol, p.Conf.addr)

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
			break
		}

		ctx := NewContext(conn, s)
		s.core.mds.run(ctx)
	}
}
