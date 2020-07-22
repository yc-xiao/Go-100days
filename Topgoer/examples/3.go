package main

import "fmt"

/*
	new() -> *T --> int, array, struct，　默认0值
	make() -> value --> channel, slice, map -> 进一步封装new，指定cap, len
*/

func testSlice() {
	s1 := make([]int, 5)
	s1 = append(s1, 1, 2, 3)
	fmt.Println(s1)

	s2 := make([]int, 0)
	s2 = append(s2, 1, 2, 3, 4)
	fmt.Println(s2)
}

func testMake() {
	var s3 []int
	if s3 == nil {
		fmt.Printf("s1 is nil --> %#v \n", s3)
		// s3[0] = 1 -> panic -> len
	} else {
		fmt.Printf("s1 is not nil --> %#v \n", s3)
	}

	s4 := make([]int, 3)
	if s4 == nil {
		fmt.Printf("s2 is nil --> %#v \n", s4)
	} else {
		fmt.Printf("s2 is not nill --> %#v \n", s4) // []int{0, 0, 0}
	}
}

func main() {
	testSlice()
	testMake()
}
