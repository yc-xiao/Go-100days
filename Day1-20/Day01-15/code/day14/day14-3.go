package day14

import "fmt"

type Interface2 interface {
	info()
	sayName()
}

type persion struct {
	name   string
	gender bool
	age    int
}

type student struct {
	persion
}

func (p persion) sayName() {
	fmt.Println("他的名字是-> ", p.name)
}

func (p persion) info() {
	fmt.Println("这是一个人类")
}

func (s student) info() {
	fmt.Println("这是一个学生")
}

func New(name string) persion {
	// 通过对外公开一个　persion　的自定义构造函数，保证　persion属性　不为空
	p := persion{name: name}
	return p
}

func Pmain14_3() {
	//OPP
	test_stuct_func()
	fmt.Println()
	test_stuct_inherit()
	fmt.Println()
	test_stuct_multi()
}

func test_stuct_func() {
	// 测试结构体函数，并限制结构体
	p := new(persion)
	p.sayName()
	p1 := New("小明")
	p1.sayName()
}

func test_stuct_inherit() {
	// 测试继承性，通过组合实现
	s := student{persion{"小明同学", true, 16}}
	s.sayName()
}

func test_stuct_multi() {
	// 测试多态性，多态由接口实现
	p := persion{name: "小明", age: 20, gender: true}
	s := student{persion{"小红", false, 18}}

	arrInterfaces := [2]Interface2{p, s}
	for _, i := range arrInterfaces {
		i.info()
		i.sayName()
	}
}
