package main

import (
	"fmt"
	"log"
	"net/rpc"
)

// 只需要属性跟Obj一样，发送对象后，会将对应值赋予Obj并返回结果
type Obj2 struct {
	A, B int
}

func main() {
	// 1.注册
	conn, err := rpc.DialHTTP("tcp", ":9099")
	if err != nil {
		log.Fatal(err)
	}
	// 2.调用
	ret := 0
	err2 := conn.Call("Obj.Add", Obj2{50, 100}, &ret)
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println("相加：", ret)
	err3 := conn.Call("Obj.Reduce", Obj2{50, 100}, &ret)
	if err3 != nil {
		log.Fatal(err3)
	}
	fmt.Println("相差：", ret)
}
