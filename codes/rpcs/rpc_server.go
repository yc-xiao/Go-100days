package main

import (
	"net/http"
	"net/rpc"
)

type Obj struct {
	A, B int
}

func (o *Obj) Add(o1 Obj, ret *int) error {
	*ret = o1.A + o1.B
	return nil
}

func (o *Obj) Reduce(o1 Obj, ret *int) error {
	*ret = o1.A - o1.B
	return nil
}

func main() {
	o := new(Obj)
	// 1. 注册
	// 2. 添加到路由
	// 3. 启动http
	rpc.Register(o)
	rpc.HandleHTTP()
	http.ListenAndServe(":9099", nil)
}
