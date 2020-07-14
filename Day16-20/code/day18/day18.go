package day18

import (
	"fmt"
	"runtime"
)

func write(cc chan int, n int) {
	for i := 1; i <= n; i++ {
		cc <- i
		fmt.Println(i, cc)
	}
}

func read(cc chan int, n int) {
	for {
		if v, ok := <-cc; ok {
			fmt.Println(v, cc)
			// 关闭通道
			if v == n {
				close(cc)
			}
		} else {
			break
		}
	}
}

func test_chan1() {
	// 缓存双向管道
	n := 5
	cc := make(chan int, n)
	write(cc, n)
	read(cc, n)
}

func test_chan2() {
	cc := make(chan int)

	// 单向管道做形参，限制管道
	go func(cc chan<- int) {
		for i := 0; ; i++ {
			cc <- i
		}
	}(cc)
	go func(cc <-chan int) {
		for {
			fmt.Println(<-cc)
		}
	}(cc)

	// 让一下时间片
	runtime.Gosched()
	fmt.Println("close")
}

func test_chan3() {
	n := 3
	cw := make(chan<- int, n)
	for i := 0; i < n+1; i++ {
		cw <- i // n+1 阻塞，没有接受者， 死锁
		fmt.Println(cw)
	}
	close(cw) // 读管道不能关闭
}

func Pmain18() {
	test_chan1()
	test_chan2()
	//test_chan3()
}
