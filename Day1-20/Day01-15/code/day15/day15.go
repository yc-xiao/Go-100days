package day15

import (
	"errors"
	"fmt"
	"strconv"
)

func check_age(age int) error {
	if age >= 18 {
		return nil
	}
	return errors.New("未成年要好好学习!")
}

func test1() {
	err := check_age(16)
	if err != nil {
		fmt.Println(err)
	}
}

type Persion struct {
	name string
	age  int
}

type NameError struct {
	Persion
}
type AgeError struct {
	Persion
}

func (p NameError) Error() string {
	// 名称不能以数字开头
	return "name error " + p.name
}

func (p *AgeError) Error() string {
	// 年龄要大于17
	return "age error " + strconv.Itoa(p.age)
}

func check_persion(p *Persion) error {
	// 检测用户是否规范
	if p.age < 18 {
		return &AgeError{*p}
	} else if len(p.name) < 3 {
		return NameError{*p}
	} else {
		return errors.New("通过New生成一个错误!")
	}
}

func test2() {
	// 自定义错误
	p := Persion{"aa", 18}
	err := check_persion(&p)
	fmt.Println(err)
	if ins, ok := err.(*AgeError); ok {
		fmt.Println(ins.age)
	} else if ins, ok := err.(NameError); ok {
		fmt.Println(ins.name)
	}
}

func test3() {
	// panic 与 recover
	n := 10
	defer func() {
		if msg := recover(); msg != nil {
			fmt.Println(msg)
			fmt.Println("想跑，没门")
		} else {
			fmt.Println("看起来风平浪静")
		}
	}()
	defer func() {
		fmt.Println("我就看看!")
	}()
	if n < 100 {
		panic("n小于100")
	}
	defer func() {
		fmt.Println("刚刚的n是大于100的")
	}()
}

func Pmain15() {
	test3()
}
