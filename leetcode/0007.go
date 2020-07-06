package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	var num int
	for x!=0{
		num = num*10 + x%10
		x/=10
	}
	if math.MaxInt32 < num || math.MinInt32 > num{
		return 0
	}
	return num
}

func main() {
	num:=123456
	num=reverse(num)
	fmt.Println(num)
}
