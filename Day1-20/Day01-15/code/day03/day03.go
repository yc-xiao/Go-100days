package day03

import (
	"fmt"
)

func Pmain3() {
	var a uint8 = 1
	var b byte = 2
	var c = a + b
	var d = int(a)
	var e = string(a)
	// 类型转换
	var f = a + uint8(d)
	fmt.Println(c, d, e, f)
	fmt.Println()

	v3 := '中' // go没有字符，用int存编码表对应数值
	v4 := "中" // 字符串 bytes 集合 uint8，utf-8一个汉字对应3个字节

	fmt.Println(string(v3) == v4)
	fmt.Printf("%T,%d,%c,%q\n", v3, v3, v3, v3)
	fmt.Printf("%T,%d,%c,%q\n", v4[0], v4[0], v4[0], v4)
	fmt.Println()
	for i, v := range v4 {
		if v == '中' {
			fmt.Println("no")
		}
		fmt.Printf("%T, %d, %p, %v \n", i, i, i, i)
		fmt.Printf("%T, %d, %p, %v \n", v, v, v, v)
	}
}
