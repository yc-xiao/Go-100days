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

func (o *Obj) Add2(ns []int, ret *int) error {
	for _, n := range ns {
		*ret += n
	}
	return nil
}

func main() {
	// 1. 注册
	// 2. 添加到路由，进行rpc处理
	// 3. 启动http
	// ps 只接收两个参数，第二个参数是固定返回值
	rpc.Register(new(Obj))
	rpc.HandleHTTP()
	http.ListenAndServe(":9099", nil)
}
