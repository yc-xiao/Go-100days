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
	rpc.Register(o)
	rpc.HandleHTTP()
	http.ListenAndServe(":9099", nil)
}
