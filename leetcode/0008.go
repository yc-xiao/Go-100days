package main

import (
	"fmt"
	"math"
)

func myAtoi(str string) int {
	var num int
	var flag bool
	for _, c := range str{
		if c == ' ' && num==0{
			continue
		}
		if c == '-' || c == '+'{
			if flag{
				return 0
			}else {
				flag=true
				continue
			}
		}
		if '0' <= c && '9' >= c{
			num = num*10 + int(c) - 48
		}else {
			break
		}
	}
	if flag{
		num*=-1
	}
	if num > math.MaxInt32{
		return math.MaxInt32
	}
	if num < math.MinInt32{
		return math.MinInt32
	}
	return num
}

func main() {
	num:=myAtoi("  +42 3asd")
	fmt.Println(num)
}
