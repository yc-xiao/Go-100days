package day05

import (
	"fmt"
	"math"
)

func Pmain5() {
	//练习1 58-23
	for i := 58; i > 22; i-- {
		fmt.Print(i, " ")
	}
	fmt.Println()

	//练习2 1-100的和
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)
	fmt.Println()

	//练习3：打印1-100内，能够被3整除，但是不能被5整除的数字，统计被打印的数字的个数，每行打印5个
	n := 0
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 != 0 {
			n++
			fmt.Print(i, " ")
			if n == 5 {
				n = 0
				fmt.Println()
			}
		}
	}
	fmt.Println()

	/*
		题目二：								i	j
		1X1=1								1	1
		2X1=2 2X2=4							2	1,2
		3X1=3 3X2=6 3X3=9					3	1,2,3
		...									...
		9X1=9 9X2=18 9X3=27...9X9=81		9	1,2,3,...9
	*/
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%dx%d=%d ", i, j, i*j)
		}
		fmt.Println()
	}

	// 练习4: 打印2-100内的素数(只能被1和本身整除)
	for i := 2; i <= 100; i++ {
		n := 0
		for j := 2; j <= int(math.Sqrt(float64(i))); j++ {
			if i%j == 0 {
				n += 1
				break
			}
		}
		if n == 0 {
			fmt.Print(i, " ")
		}
	}

}
