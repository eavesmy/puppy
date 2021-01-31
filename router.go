package puppy

import (
	"github.com/teambition/trie-mux"
	"net/http"
	"strings"
)

// Use router need a middleware to deal net packet before this work.

type Router struct {
	mds        []Middleware
	middleware Middleware
	trie       *trie.Trie
	Root       string
	// rpcMethods []rpc_method
}

func NewRouter(confs ...RouterConf) *Router {

	conf := DefaultRouterConf()

	if len(confs) > 0 {
		conf = confs[0]
	}

	if conf.Root != "" && conf.Root[0] != '/' {
		conf.Root += "/"
	}

	return &Router{mds: make([]Middleware, 0), trie: trie.New(trie.Options{
		IgnoreCase:            conf.IgnoreCase,
		FixedPathRedirect:     conf.FixedPathRedirect,
		TrailingSlashRedirect: conf.TrailingSlashRedirect,
	}), Root: conf.Root}

}

// rpc 调用的只能是一个带有返回值的函数
func (r *Router) Rpc(pattern string, handlers ...Middleware) *Router {
	/*
		if r.rpcMethods == nil {
			r.rpcMethods = make([]rpc_method, 0)
		}
	*/

	if string([]rune(pattern)[0]) != "/" {
		pattern = "/" + pattern
	}

	// r.rpcMethods = append(r.rpcMethods, rpc_pack_methods(pattern, handlers...))

	return r.Handle("RPC", pattern, handlers...)
}

/*
func (r *Router) encodeRpcMethods() []rpc_method {
	return r.rpcMethods
}
*/

func (r *Router) Get(pattern string, handlers ...Middleware) *Router {
	return r.Handle(http.MethodGet, pattern, handlers...)
}

func (r *Router) Post(pattern string, handlers ...Middleware) *Router {
	return r.Handle(http.MethodPost, pattern, handlers...)
}

func (r *Router) Put(pattern string, handlers ...Middleware) *Router {
	return r.Handle(http.MethodPut, pattern, handlers...)

}

func (r *Router) Delete(pattern string, handlers ...Middleware) *Router {
	return r.Handle(http.MethodDelete, pattern, handlers...)
}

func (r *Router) Options(pattern string, handlers ...Middleware) *Router {
	return r.Handle(http.MethodOptions, pattern, handlers...)
}

func (r *Router) Handle(method, pattern string, handlers ...Middleware) *Router {
	if method == "" {
		panic("required method.")
	}

	if string([]rune(pattern)[0]) != "/" {
		pattern = "/" + pattern
	}

	pattern = r.Root + pattern

	r.trie.Define(pattern).Handle(strings.ToUpper(method), compose(handlers...))
	return r
}

func (r *Router) Serve(ctx *Context) (err error) {

	if r.middleware != nil {
		if err = compose(r.middleware)(ctx); err != nil {
			return
		}
	}

	var handler Middleware

	if ctx.Path == "" {
		ctx.Path = "/"
	}

	if string([]rune(pattern)[0]) != "/" {
		ctx.Path = "/" + ctx.Path
	}

	matched := r.trie.Match(ctx.Path)
	// 找不到对应节点

	if matched.Node == nil {
		return ctx.Text("invalid method", 405)
	}

	ok := false
	if handler, ok = matched.Node.GetHandler(ctx.Method).(Middleware); !ok {

		return
	}

	return handler(ctx)
}

// Router will load this method before router serve.
func (r *Router) Use(handle Middleware) *Router {
	r.middleware = compose(handle)
	return r
}

func compose(handlers ...Middleware) Middleware {

	if len(handlers) > 1 {
		return middlewares(handlers).run
	} else if len(handlers) == 0 {
		return func(ctx *Context) error { return nil }
	} else {
		return handlers[0]
	}
}
