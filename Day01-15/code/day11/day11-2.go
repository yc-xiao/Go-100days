package day11

import "fmt"

// 包下所有文件的全局变量名，大小写都定义一次
const n = 1000
const N = 100

func init() {
	fmt.Println("day11-2 init")
}

func Pmain12() {
	fmt.Print(N)
}
