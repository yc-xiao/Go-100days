package day09

import (
	"fmt"
	"strconv"
	"strings"
)

func Pmain8() {
	var is bool
	var index, count int
	s1 := "helloc 中"

	// strings.Contains(string, substr)
	is = strings.Contains(s1, "中")
	fmt.Println(is)
	is = strings.Contains(s1, "中1")
	fmt.Println(is)

	// strings.Index
	index = strings.Index(s1, "中")
	fmt.Println(index)
	index = strings.Index(s1, "1中")
	fmt.Println(index)

	count = strings.Count(s1, "l")
	fmt.Println(count)

	s2 := strings.Join([]string{"h", "e", "l", "l", "o"}, "")
	fmt.Println(s2)

	s3 := strings.Split(s2, "")
	fmt.Println(s3)

	// 字符串<==>数字
	s, n := "10", 10
	sn, _ := strconv.ParseInt(s, 10, 32)
	fmt.Printf("%T, %v, %T, %v\n", s, s,  sn, sn)
	ns := strconv.FormatInt(int64(n), 10)
	fmt.Printf("%T, %v, %T, %v", n, n,  ns, ns)

}