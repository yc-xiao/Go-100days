package day17

import (
	"fmt"
	"sync"
	"time"
)

func reduce(num *int) {
	for *num > 0 {
		time.Sleep(time.Millisecond / 10)
		*num--
	}
}
func test_go3() {
	var num = 100
	fmt.Println(num)
	for i := 0; i < 10; i++ {
		go reduce(&num)
	}
	time.Sleep(time.Second / 10)
	fmt.Println(num)
}

func reduce2(num *int, wg *sync.WaitGroup, mutex *sync.Mutex) {
	defer wg.Done()
	for {
		mutex.Lock()
		if *num > 0 {
			time.Sleep(time.Millisecond / 10)
			*num--
		} else {
			break
		}
		mutex.Unlock()
	}
}

func test_go4() {
	// 定义一个等待组
	// 定义一个变量锁
	var wg sync.WaitGroup
	var mutex sync.Mutex
	var num = 100
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go reduce2(&num, &wg, &mutex)
	}
	time.Sleep(time.Second)
	fmt.Println(num)
}

func Pmain17_2() {
	//test_go3()
	test_go4()
}
