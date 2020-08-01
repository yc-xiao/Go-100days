package main

import (
	"fmt"
	"sync"
	"time"
)

var mu sync.Mutex
var muRW sync.RWMutex

const T = time.Second
const GoNum = 10

func test1(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	go func(i int) {
		mu.Lock()
		for j := 0; j < 10; j++ {
			fmt.Printf("%d -> %d  ", i, j)
		}
		fmt.Println()
		mu.Unlock()
	}(i)
}

// 加锁需要在 gorountine 内加
func test2(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	mu.Lock()
	go func(i int) {
		for j := 0; j < 10; j++ {
			fmt.Printf("%d -> %d  ", i, j)
		}
		fmt.Println()
	}(i)
	mu.Unlock()
}

func test3(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	go func(i int) {
		muRW.Lock()
		for j := 0; j < 10; j++ {
			fmt.Printf("%d -> %d  ", i, j)
		}
		muRW.Unlock()
		fmt.Println()
	}(i)
}

func test4(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	go func(i int) {
		muRW.RLock()
		for j := 0; j < 10; j++ {
			fmt.Printf("%d -> %d  ", i, j)
		}
		muRW.RUnlock()
		fmt.Println()
	}(i)
}

func test(f func(i int, wg *sync.WaitGroup), funName string) {
	var wg sync.WaitGroup
	wg.Add(GoNum)
	for i := 0; i < GoNum; i++ {
		f(i, &wg)
	}
	wg.Wait()
	fmt.Println()

	// 并发无法控制顺序，wg无法控制多个并发的顺序
	// 此处没有使用go协程， 用休眠等待
	time.Sleep(T * 3)
}

func main() {
	test(test1, "test1")
	test(test2, "test2")
	test(test3, "test3")
	test(test4, "test4")
}
