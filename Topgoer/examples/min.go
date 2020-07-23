package examples

import "fmt"

// Channel

// Ptr
var Func31 = Func31C
var Func30 = Func30C
var Func29 = Func29C

// 闭包
var Func20 = Func20C
var Func19 = Func19C
var Func18 = Func18C
var Func17 = Func17C

func Func41() {
	var x = []int{2: 2, 3, 0: 1}
	fmt.Println(x) // slice在初始化时可以指定值
}

func Func37_1() {
	//const i = 100 // &i -> err
	var i = 100
	fmt.Println(&i, i)
}

func Func37() {
	/*
		A. 给一个 nil channel 发送数据，造成永远阻塞
		B. 从一个 nil channel 接收数据，造成永远阻塞
		C. 给一个已经关闭的 channel 发送数据，引起 panic
		D. 从一个已经关闭的 channel 接收数据，如果缓冲区中为空，则返回一个零值
	*/
	c := make(chan int)
	A := func(c chan int) {
		c <- 1
	}
	B := func(c chan int) {
		data := <-c
		fmt.Println(data)
	}
	C := func(c chan int) {
		defer func() {
			msg := recover()
			if msg != nil {
				fmt.Println("它慌了它慌了，它就要跳墙走了!!!!")
			}
		}()
		close(c)
		c <- 1
	}
	D := func(c chan int) {
		close(c)
		data := <-c
		fmt.Println(data)
	}
	fmt.Println(A, B, C, D)
	C(c)
}

func Func27() {
	type Math struct {
		x, y int
	}
	var m = map[string]Math{"foo": Math{2, 3}}
	// m["foo"].x = 4 -> err 无法获取值
	v := m["foo"]
	v.x = 4
	m["foo"] = v
	fmt.Println(m["foo"])

	var m2 = map[string]*Math{"foo": &Math{2, 3}}
	m2["foo"].x = 4
	fmt.Println(m2["foo"])
}

func Func26() {
	const (
		a = iota
		b = iota
	)
	const (
		name = "name"
		c    = iota
		d    = iota
	)
	fmt.Println(a) // 0
	fmt.Println(b) // 1
	fmt.Println(c) // 1
	fmt.Println(d) // 2
}

func Func23() {
	s1 := []int{1, 2, 3}
	s2 := s1[1:]
	s2[1] = 4
	fmt.Println(s1) // 1, 2, 4
	s2 = append(s2, 5, 6, 7)
	fmt.Println(s1)          // 1, 2, 4
	fmt.Println(s2, cap(s2)) // 2, 4, 5, 6, 7, cap(s2)->6
}

func Func16_2() {
	// var m map[string]int
	m := make(map[string]int) // 必须实例化，不然为空无法使用
	m["a"] = 1
	fmt.Println(m)
	if v, ok := m["b"]; ok { //B
		fmt.Println(v)
	}
}

func Func16() {
	s := [3]int{1, 2, 3}
	a := s[:0]
	b := s[:2]
	c := s[1:2:cap(s)]
	fmt.Println(len(a), cap(a)) // 0, 3 => s[i:j] => j-i, len(s)-i
	fmt.Println(len(b), cap(b)) // 2, 3 => s[i:j] => j-i, len(s)-i
	fmt.Println(len(c), cap(c)) // 1, 2 => s[i:j:k] => j-i, k-i
}

func Func15_2() {
	i := 65
	fmt.Println(string(i)) // -> A
}

func Func15() {
	var s1 []int
	var s2 = []int{}
	if s1 == nil && s2 == nil {
		fmt.Println("s1,s2 == nil")
	} else if s1 != nil && s2 != nil {
		fmt.Println("s1, s2 != nil")
	} else if s1 == nil && s2 != nil {
		fmt.Println("s1==nil, s2!=nil")
	} else if s1 != nil && s2 == nil {
		fmt.Println("s1!=nil, s2==nil")
	} else {
		fmt.Println("what ???")
	}
}

func Func14_2() {
	add := func(args ...int) int {
		sum := 0
		for _, arg := range args {
			sum += arg
		}
		return sum
	}
	add(1, 2, 3)
	add([]int{1, 2, 3}...)
}

func Func14() {
	str := "hello"
	//str[0] = 'x' // err str常量
	fmt.Println(str)
}

func Func10() {
	// slice[::]
	a := [5]int{1, 2, 3, 4, 5}
	t := a[3:4:4]
	fmt.Println(t[0])
}

func Func9() {
	// 可变量
	hello := func(num ...int) {
		num[0] = 18
	}

	i := []int{5, 6, 7}
	hello(i...)
	fmt.Println(i[0])
}

func Func7() {
	// iota
	const (
		x = iota
		_
		y
		z = "zz"
		k
		p = iota
	)
	fmt.Println(x, y, z, k, p)
}

func Func6() {
	type MyInt1 int
	type MyInt2 = int
	var i int = 0
	//var i1 MyInt1 = i
	var i2 MyInt2 = i
	_ = i2
	//fmt.Println(i1, i2)
}
