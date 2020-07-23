package examples

import "fmt"

/*
方法必须在同一个包内定义
func (i int) PrintInt ()  {
	fmt.Println(i)
}
var i int = 1
i.PrintInt()
*/

type int2 int // 定义一种新的类型
func (i int2) PrintInt() {
	fmt.Println(i)
}

func test25_1() {
	var i int2 = 1 // var i int = 1
	i.PrintInt()
}

func Func25() {
	test25_1()
}
