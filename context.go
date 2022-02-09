/*
# File Name: ./context.go
# Author : eavesmy
# Email:eavesmy@gmail.com
# Created Time: 三  2/ 9 18:28:44 2022
*/

package puppy

import (
	"context"
	"encoding/json"
	"net"
)

type SettingPreSend func(interface{}) []byte

func DefaultSettingPreSend(body interface{}) []byte {
	b, _ := json.Marshal(body)
	return b
}

type Context struct {
	Context *context.Context
	Socket  net.Conn

	FnPreSend SettingPreSend
}

func (c *Context) Send(body interface{}) {
	b := c.FnPreSend(body)
	c.Socket.Write(b)
}

// Before send msg,you can convert your own struct to bytes.
func (c *Context) SetPreSendCfg(fn SettingPreSend) {
	c.FnPreSend = fn
}
