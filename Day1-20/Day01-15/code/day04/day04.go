package day04

import "fmt"

func Pmain4() {
	if n := 1; n < 0 {
		fmt.Printf("%d < 0 \n", n)
	} else if n == 1 {
		fmt.Printf("%d = 1 \n", n)
	} else {
		fmt.Printf("%d > 0 and %d != 1 \n", n, n)
	}

	n := 3 // 重新声明并赋值n
	if n < 0 {
		fmt.Printf("%d < 0 \n", n)
	} else if n == 1 {
		fmt.Printf("%d = 1 \n", n)
	} else {
		fmt.Printf("%d > 0 and %d != 1 \n", n, n)
	}

	switch n1 := 3; n1 {
	case 3:
		fmt.Printf("n1的值为3\n")
	case 4, 5, 6:
		fmt.Printf("n1的值为4-5-6\n")
	}

	switch {
	case n > 13:
		fmt.Printf("n的值为3\n")
	default:
		fmt.Printf("n的值为<=13\n")
	}
}
