package day01

import "fmt"

func SayHello() {
	PrintC() // 包内直接引用
	fmt.Println("hello world!")
}