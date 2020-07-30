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
	/*
		PS:
			1) 方法类型（T）是导出的（首字母大写）
			2) 方法名（MethodName）是导出的
			3) 方法有2个参数(argType T1, replyType *T2)，均为导出/内置类型
			4) 方法的第2个参数一个指针(replyType *T2)
			5) 方法的返回值类型是 error
	*/
	rpc.Register(new(Obj))            // 1. 注册
	rpc.HandleHTTP()                  // 2. 添加到路由，进行rpc处理
	http.ListenAndServe(":9099", nil) // 3. 启动http
}
