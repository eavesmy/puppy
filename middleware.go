package puppy

import (
	"fmt"
)

type Handler interface {
	Serve(ctx *Context) error
}

// Middleware get ctx param. After data handled return error.
type Middleware func(*Context) error

type middlewares []Middleware

// call middlewares.
func (m middlewares) run(ctx *Context) (err error) {

	for _, fn := range m {
		if err := fn(ctx); err != nil {
			fmt.Println(err)
		}
	}
	return
}
