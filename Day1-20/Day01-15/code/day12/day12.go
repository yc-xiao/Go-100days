package day12

import "fmt"

func testp() {
	// 指针 变量 -> 变量地址 -> 变量存储值
	n := 100
	fmt.Printf("%p, %d\n", &n, n)
	np := &n
	fmt.Printf("%p, %d\n", np, *np)
	*np += 1
	fmt.Printf("%d, %d\n", n, *np)
}

func testpp() {
	n := 100
	np := &n
	npp := &np
	**npp += 1
	fmt.Printf("%d, %p, %p, %d, %p, %d, \n", n, &n, np, *np, &np, **npp)
}

func modifyByvalue(arr [3]int) {
	arr[0] = 100
}

func modifyBypoint(arr *[3]int) {
	arr[0] = 200
}

func modifyByslice(arr []int) {
	arr[0] = 300
}

func test_arr() {
	// 一般用切片代替数组指针
	var arr [3]int
	arr = [3]int{1, 2, 3}
	modifyByvalue(arr)
	fmt.Println(arr)

	modifyBypoint(&arr)
	fmt.Println(arr)

	arr = [3]int{1, 2, 3}
	modifyByslice(arr[:])
	fmt.Println(arr)
}

func Pmain12() {
	testp()
	testpp()
	test_arr()
}
