package puppy

import (
	"time"
)

const (
	DEFAULT_MAXRECONNECTCOUNT int           = 5
	DEFAULT_RECONNECTTIMEOUT  time.Duration = time.Second * 5
	DEFAULT_BUFFERSIZE        int           = 512
)

type Puppy struct {
	Server *Server
	Conf   Conf
	mds    middlewares
	router *Router
}

// Init socket server.
func New(confs ...Conf) *Puppy {

	var conf Conf
	def := DefaultConf()

	if len(confs) > 0 {
		conf = confs[0]
		if conf.MaxReconnectCount == 0 {
			conf.MaxReconnectCount = def.MaxReconnectCount
		}
		if conf.ReconnectTimeout == 0 {
			conf.ReconnectTimeout = def.ReconnectTimeout
		}
		if conf.BufferSize == 0 {
			conf.BufferSize = def.BufferSize
		}
	}

	return &Puppy{Conf: conf, Server: &Server{}}
}

// Load middleware
func (p *Puppy) Use(m Middleware) *Puppy {
	p.mds = append(p.mds, m)
	return p
}

func (p *Puppy) UseHandle(m Handler) *Puppy {
	p.mds = append(p.mds, m.Serve)
	return p
}

func (p *Puppy) Run(h_ps ...string) error {

	if len(h_ps) > 0 {
		p.Conf.addr = h_ps[0]
	}

	return p.Server.Listen(p)
}

func (p *Puppy) Rpc(pattern string, fn func(*Context, interface{}) interface{}) {

	return
}
