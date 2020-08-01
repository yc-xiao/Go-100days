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

func test(f func(i int, wg *sync.WaitGroup)) {
	var wg sync.WaitGroup
	wg.Add(GoNum)
	for i := 0; i < GoNum; i++ {
		//f(i, &wg)  // 会立马退出，不等待子go
		go f(i, &wg) // 并发执行，执行完解锁，等待子go
	}
	wg.Wait()
	fmt.Println()
}

func test5() {
	var wg sync.WaitGroup
	wg.Add(GoNum)
	for i := 0; i < GoNum; i++ {
		// 会立马退出，不等待子go
		func(i int) {
			defer wg.Done()
			go func() {
				fmt.Println(i)
			}()
		}(i)
	}
	wg.Wait()

	fmt.Println("--------------")
	wg.Add(GoNum)
	for i := 0; i < GoNum; i++ {
		// 并发执行，执行完解锁，等待子go
		go func(i int) {
			defer wg.Done()
			go func() {
				fmt.Println(i)
			}()
		}(i)
	}
	wg.Wait()
	fmt.Println("--------------")
}

func main() {
	//fs := [...]func(i int, wg *sync.WaitGroup){test1, test2, test3, test4}
	type ff = func(i int, wg *sync.WaitGroup)
	fs := [...]ff{test1, test2, test3, test4}

	var wg sync.WaitGroup
	wg.Add(len(fs))
	for _, f := range fs {
		test(f)
		wg.Done() // 单行执行，执行完释放
	}
	wg.Wait()
	fmt.Println("test1-4")
	//test5()
	fmt.Println("main over")
}
