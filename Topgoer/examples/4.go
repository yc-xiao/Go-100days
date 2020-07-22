package main

import "fmt"

func test1() {
	list := new([]int)
	*list = append(*list, 0) // err -> list = append(list, 0) // list是指针没有append方法
	fmt.Println(list)
}

func test2() {
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5}
	s1 = append(s1, s2...) // err -> s1 = append(s1, s2)
	fmt.Println(s1)
}

func test3() {
	var (
		size    = 1024 // err -> size := 1024 与 var 冲突
		maxSize = size * 2
	)
	fmt.Println(size, maxSize)
}

func main() {
	test1()
	test2()
	test3()
}
