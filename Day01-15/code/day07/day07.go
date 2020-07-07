package day07

import "fmt"

func test1() {
	// 切片引用数组，修改会影响数组
	arr := [3]int{}
	a1 := arr[:]
	arr[0] = 1
	a1[1] = 2
	fmt.Println(arr, a1)
}

func test2() {
	// append返回一个新对象
	arr := [3]int{1}
	a1 := arr[:]
	a2 := append(a1, 3)
	a2[0] = 2
	fmt.Println(arr, a1, a2)
	fmt.Printf("%p, %p, %p, \n", arr, a1, a2)
}

func Pmain7() {
	test1()
	test2()
}
