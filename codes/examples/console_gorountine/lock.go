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

func test1(wg *sync.WaitGroup) {
	// 测试lock
	wg.Add(GoNum)
	for i := 0; i < GoNum; i++ {
		// mu.Lock 加在此处错误， go func() 开启一个新的协程
		go func(i int) {
			defer wg.Done() // 在子协程内释放
			mu.Lock()
			for j := 0; j < 10; j++ {
				fmt.Printf("%d -> %d  ", i, j)
			}
			fmt.Println()
			mu.Unlock()
		}(i)
	}
}


func test2(wg *sync.WaitGroup) {
	// 测试RWlock
	wg.Add(GoNum)
	for i := 0; i < GoNum; i++ {
		go func(i int) {
			defer wg.Done() // 在子协程内释放
			muRW.RLock() // 读锁可以重复读, muRW.Lock == mu.Lock --> 写锁
			for j := 0; j < 10; j++ {
				fmt.Printf("%d -> %d  ", i, j)
			}
			fmt.Println()
			muRW.RUnlock()
		}(i)
	}
}

func test3(wg *sync.WaitGroup) {
	fp := func(n int) {
		time.Sleep(T)
		fmt.Println(n)
	}
	wg.Add(2)
	func() {
		defer wg.Done()
		fp(1)
	}()
	func() {
		defer wg.Done() // 秒释放，不会等待fp(2)
		go fp(2)
	}()
	wg.Wait()
}

func test(t func(wg *sync.WaitGroup)) {
	var wg sync.WaitGroup
	t(&wg)
	wg.Wait()
	fmt.Println()
}

func main() {
	test(test1)
	test(test2)
	test(test3)
}
