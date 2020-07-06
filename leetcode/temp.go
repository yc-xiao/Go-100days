package main

import (
	"fmt"
	"math"
	"strconv"
)

func get_max() int {
	max := 1
	for i := 1; i <= 31; i++ {
		max *= 2
	}
	return max
}

func reverse(x int) int {
	var reverse_x int
	for x != 0 {
		reverse_x = reverse_x*10 + x%10
		x /= 10
	}
	if reverse_x >= get_max() || reverse_x <= math.MinInt32 {
		return 0
	}
	return reverse_x
}

func is_ok(num, reverse_num int) bool {
	if reverse(num) != reverse_num {
		fmt.Println(num, reverse_num)
		defer func() {
			if error := recover(); error != nil {
				//fmt.Println(error)
			}
		}()
		error := strconv.Itoa(num) + "!=" + strconv.Itoa(reverse_num)
		panic(error)
	}
	return true
}

func main() {
	map_temp := map[int]int{123: 321, -123: -321, 120: 21, 1534236469: 0}
	for num, reverse_num := range map_temp {
		ok := is_ok(num, reverse_num)
		fmt.Println(ok)
	}
}
