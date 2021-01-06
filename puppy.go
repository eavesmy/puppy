package puppy

import (
	"fmt"
	"github.com/teambition/trie-mux"
)

type Puppy struct {
	Server *Server
	Conf   Conf
	mds    middlewares
	router *Router
	trie   *trie.Trie
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

	return &Puppy{Conf: conf, Server: &Server{}, trie: trie.New(trie.Options{
		IgnoreCase:            conf.IngoreCase,
		FixedPathRedirect:     false,
		TrailingSlashRedirect: false,
	})}
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

func (p *Puppy) MarshalByHand() {

}

func (p *Puppy) ParseByHand() {

}

func (p *Puppy) Run(h_ps ...string) error {

	if len(h_ps) > 0 {
		p.Conf.addr = h_ps[0]
	}

	return p.Server.Listen(p)
}

// func(*Context,interface{}) interface{}

func (p *Puppy) Rpc(pattern string, fn Rpc) {

	method, err := MarshalRpcMethod(pattern, fn)

	if err != nil {
		panic("Invalid rpc method: " + err.Error())
	}

	fmt.Println(method)

	return
}
