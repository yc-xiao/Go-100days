package day10

import "fmt"

//定义函数类型，用于参数传递
type funName func(inputN int) (outN int)

func Pmain10() {
	n := 20

	// 匿名函数
	// defer 延时函数，闭包
	defer func(n int) {
		fmt.Println(n)
	}(n)

	n += 10

	funA := func(n int) {
		fmt.Println(n)
	}
	defer funA(n)

}
