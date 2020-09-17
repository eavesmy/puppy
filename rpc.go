package puppy

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"reflect"
)

type Rpc interface{}

type rpcMethod struct {
	Name    string
	Args    [][]byte
	Returns [][]byte
}

/*
	RPC 方法参数与返回值封装
	提供给本节点作为暴露 rpc routes 使用
*/

func MarshalRpcMethod(pattern string, fn interface{}) (method *rpcMethod, err error) {

	t := reflect.TypeOf(fn)

	if t.Kind().String() != "func" {
		return nil, errors.New("Invalid type. Required func")
	}

	args_num := t.NumIn()

	if args_num < 1 {
		return nil, errors.New("Invalid args num. First argument must be *Context")
	}

	method = &rpcMethod{Name: pattern, Args: [][]byte{}, Returns: [][]byte{}}

	for i := 1; i < args_num; i++ {
		v := reflect.New(t.In(i))
		arg := EncodeByGob(v.Interface())
		method.Args = append(method.Args, arg)
	}

	for i := 0; i < t.NumOut(); i++ {
		v := reflect.New(t.Out(i))
		r := EncodeByGob(v.Interface())
		method.Returns = append(method.Returns, r)
	}

	return
}

func EncodeByGob(i interface{}) []byte {

	var buffer bytes.Buffer

	enc := gob.NewEncoder(&buffer)

	if err := enc.Encode(&i); err != nil {
		fmt.Println(err)
		return []byte{}
	}

	return buffer.Bytes()
}
