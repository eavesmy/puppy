/*
# File Name: main.go
# Author : eavesmy
# Email:eavesmy@gmail.com
# Created Time: 三  2/ 9 17:25:53 2022
*/

package main

import (
	"github.com/eavesmy/puppy"
)

func main() {
	app := puppy.New() // create new service.

	app.UseHandler(&Handler{})

	app.Listen(":8080")
}

type Handler struct {
	// compoent.Base
	puppy.App
}

func (h *Handler) Init(app *puppy.App) {
	h.App = app
}

func (h *Handler) Entery(ctx *puppy.Context) error {
	return ctx.Send("gogo")
}
