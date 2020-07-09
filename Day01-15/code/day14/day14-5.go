package day14

import "fmt"

/*
	接口　断言 -> 获取真实类型
	接口　嵌入

*/

type A interface {
	sayA()
}
type B interface {
	sayB()
}

//　实现接口C，需要实现接口A+B
type C interface {
	A
	B
	sayC()
}
type Sa struct {
}

func (s Sa) sayA() {
	fmt.Println("sayA")
}

type Sb struct {
	Sa
}

func (s Sb) sayB() {
	fmt.Println("sayB")
}

type Sc struct {
	Sa
	Sb
}

func (s Sc) sayC() {
	fmt.Println("sayC")
}

func test_interfaceA(a A) {
	a.sayA()
}

func test_interfaceC(c C) {
	c.sayA()
	c.sayB()
	c.sayC()
}

func test1() {
	// 测试嵌套
	a := Sa{}
	c := Sc{}
	test_interfaceA(a)
	test_interfaceA(c)
	test_interfaceC(c)
}

func test_infertace(i interface{}) {
	if ins, ok := i.(Sa); ok {
		fmt.Println(ins)
	} else if ins, ok := i.([]int); ok {
		fmt.Println("切片的第二个元素为:", ins[1])
	} else {
		fmt.Println(ins)
	}
}

func test2() {
	// 测试断言
	a := Sa{}
	b := []int{1, 2, 3}
	c := "23333"
	test_infertace(a)
	test_infertace(b)
	test_infertace(c)
}

func test3() {
	//　指针
	b := Book{"xm", "helloc"}
	var i interface{} = &b
	fmt.Printf("b -> tyep=%T, addr=%p, value=%v \n", b, &b, b)
	fmt.Printf("i -> tyep=%T, addr=%p, value=%v \n", i, &i, i)

	ins := i.(*Book)
	fmt.Printf("ins -> tyep=%T, addr=%p, value=%v \n", ins, &ins, ins)

	ins.name = "xh"
	fmt.Printf("ins -> tyep=%T, addr=%p, value=%v \n", ins, &ins, ins)
	fmt.Printf("b -> tyep=%T, addr=%p, value=%v \n", b, &b, b)
	fmt.Printf("i -> tyep=%T, addr=%p, value=%v \n", i, &i, i)
}

func Pmain14_5() {
	test3()
}

type Book struct {
	name string
	info string
}
