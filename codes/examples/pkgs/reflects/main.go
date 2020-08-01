package main

import (
	"fmt"
	"reflect"
)

/*
	1.　反射是基于interface{}
	2.  interface　类似结构体，　还包含两个指针(value, type) 存储接收的变量
		i interface{
			func a()
			// *v  <- reflect.ValueOf()
			// *p  <- reflect.TypeOf()
		}
	3. Elem 基于指针对象的一套操作
*/

type Person struct {
	Name string `json:"" desc:""`
	Age  int
	Addr string
}

func handlerInterface(p interface{}) {
	pType := reflect.TypeOf(p)
	pValue := reflect.ValueOf(p)
	fmt.Println("kind是一样的，都是指变量类型(*pte/struct/int/...) --> ", pType.Kind(), pValue.Kind())
	if pType.Kind() == reflect.Ptr {
		fmt.Println("这个变量是指针，快去修改他...")
		type1 := pType.Elem()
		value := pValue.Elem()
		for i := 0; i < type1.NumField(); i++ {
			field, value := type1.Field(i), value.Field(i)
			fmt.Printf("字段 type -> %v, value -> %v \n", field, value)
			switch field.Name {
			case "Name":
				value.SetString("小黄")
			case "Age":
				value.SetInt(18)
			}
		}

	} else {
		fmt.Println("这个变量是结构体，他只能取值...")
		for i := 0; i < pType.NumField(); i++ {
			field, value := pType.Field(i), pValue.Field(i)
			fmt.Printf("字段 type -> %v, value -> %v \n", field, value)
		}
	}
	fmt.Println()
}

func main() {
	p := Person{"小明", 19, "住在马路边"}
	fmt.Printf("p -- > %v \n\n", p)
	handlerInterface(p)
	handlerInterface(&p)
	fmt.Println("p -- >", p)
}
