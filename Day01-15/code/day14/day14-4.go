package day14

import (
	"fmt"
)

// 定义结构体
type Stuct struct {
}

// 定义接口
type Interface interface {
}

// 定义函数
type Func func()

// 定义别名
type myint = int64

func Pmain14_4() {
	var n myint = 10
	fmt.Printf("%d, %T　\n", n, n)
}
