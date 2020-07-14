package day18

import (
	"fmt"
	"time"
)

func test_timer1() {
	// timer 计时器，开始计时，计时完成后会发送通知
	timer := time.NewTimer(time.Second)
	fmt.Println(timer)
	fmt.Println(2333333)
	fmt.Println(<-timer.C)
	fmt.Println(2333333)
}

func Pmain18_2() {
	test_timer1()
}
