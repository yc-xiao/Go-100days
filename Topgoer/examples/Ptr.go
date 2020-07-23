package examples

import (
	"fmt"
	"time"
)

func Func31C() {
	var m = map[string]int{
		"A": 21,
		"B": 22,
		"C": 23,
	}
	counter := 0
	// range 会复制对象，but slice, map, chann 类似指针
	for k, v := range m {
		if counter == 0 {
			delete(m, "A")
		}
		counter++
		fmt.Println(k, v)
	}
	fmt.Println("counter is ", counter)
}

func Func30C() {
	var a = [5]int{1, 2, 3, 4, 5}
	var b = a[:]
	var r [5]int
	// for i, v := range b
	for i, v := range a {
		if i == 0 {
			a[1] = 12
			a[2] = 13
		}
		r[i] = v
	}
	fmt.Println("r = ", r)
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

}

func Func29C() {
	// 正常退出
	v := []int{1, 2, 3}
	for i := range v {
		v = append(v, i)
	}
	fmt.Println("ok")

	// for i,v := range m, 变量v只创建一次，后面通过赋值给予
	var m = [...]int{1, 2, 3}
	for i, v := range m {
		go func() {
			// 没有i,v 向外找到 i,v 此时的i,v都是最大值
			fmt.Println(i, v)
		}()
	}
	time.Sleep(time.Second / 2)
	fmt.Println()

	m = [...]int{1, 2, 3}
	for i, v := range m {
		go func(i, v int) {
			fmt.Println(i, v)
		}(i, v)
	}
	time.Sleep(time.Second / 2)
}
