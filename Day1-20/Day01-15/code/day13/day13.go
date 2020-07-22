package day13

import (
	"fmt"
	"strconv"
	"strings"
)

var pstr = strings.Repeat(">", 100)

// 定义结构体
type Persion struct {
	name string
	int
}

func createPersion1(name string) Persion {
	return Persion{name: name, int: 18}
}

func createPersion2(name string) *Persion {
	return &Persion{name: name, int: 18}
}

// 结构体方法
func (p Persion) String() string {
	return "name->" + p.name
}

// 结构体方法
func (p Persion) sleep(second int) string {
	return p.name + "睡了" + strconv.FormatInt(int64(second), 10) + " S"
}

func test_init() {
	fmt.Println(pstr)
	fmt.Println("结构体初始化")
	// 初始化
	var p1 Persion
	fmt.Printf("%p, %p, %v\n", p1, &p1, p1)

	// 初始化并按顺序赋值
	p2 := Persion{"p2", 18}
	fmt.Printf("%p, %p, %v\n", p2, &p2, p2)

	// 初始化指定参数赋值
	p3 := Persion{int: 10}
	fmt.Printf("%p, %p, %v\n", p3, &p3, p3)

	// new 初始化，返回指针
	p4 := new(Persion)
	fmt.Printf("%p, %p, %v\n", p4, &p4, p4)

	// func -> 默认值
	p5 := createPersion1("p5")
	fmt.Printf("%p, %p, %v\n", p5, &p5, p5)

	// func2 -> 默认值
	p6 := createPersion2("p6")
	fmt.Printf("%p, %p, %v\n", p6, &p6, p6)
	// 结构体 . 判断是变量还是地址，　如果是变量则取地址再取值， 如果是地址则直接取值
	fmt.Println(p5.int, p6.int)
	fmt.Println()
}

func test_func() {
	fmt.Println(pstr)
	fmt.Println("结构体函数")
	xm := Persion{name: "小米"}
	fmt.Printf("%s\n", xm)
	fmt.Println(xm.sleep(2))
	xm.sleep(10)
	fmt.Println()
}

func Pmain13() {
	test_init()
	test_func()
}
