package day11

import "fmt"

func init() {
	fmt.Println("第一次初始化")
	fmt.Println(N)
}

func init() {
	fmt.Println("第二次初始化")
	fmt.Println(n)
}

/*
// 包下所有文件的全局变量名，大小写都定义一次
day11-2下已定义，可以直接用
const n = 100
const N = 1000
*/

func Pmain11() {
	// init 可以定义多次
	/*
		import (
			_ "Go-100days/Day01-15/code/day11"
		)
		调用init
			day11-2 init
			第一次初始化
			100
			第二次初始化
			1000
	*/
}
