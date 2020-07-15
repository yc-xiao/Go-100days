package day19

import (
	"fmt"
	"reflect"
)

func test3(num float64){
	fmt.Println(&num)
	n_type := reflect.TypeOf(num)
	if n_type.Kind() != reflect.Float64{
		fmt.Println("error type no float64")
		return
	}
	value := reflect.ValueOf(num)
	if value.Kind() != reflect.Float64{
		fmt.Println("error type no float64")
		return
	}
	num2 := value.Interface().(float64)
	fmt.Println(&num2)
	fmt.Println()
}

func test4(num *float64) {
	fmt.Println(num, *num)
	value := reflect.ValueOf(num)
	fmt.Println(value)
	if value.Kind() != reflect.Ptr{
		fmt.Println("error type no float64", value.Kind())
		return
	}
	v := value.Interface().(*float64)
	*v = 2.3333
	fmt.Println(v, *num)
	fmt.Println(num, *num)
}

func Pmain19_2()  {
	var num float64 = 1.23
	test3(num)
	test4(&num)
}