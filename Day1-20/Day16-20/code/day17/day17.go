package day17

import (
	"fmt"
	"runtime"
	"time"
)

func print_num(n int, name string) {
	fmt.Println()
	fmt.Print(name)
	for n >= 0 {
		fmt.Printf(" %d", n)
		n--
	}
	fmt.Println()
}

func test_go() {
	n := 10
	go print_num(n*10, "son1")
	go print_num(n*10, "son2")
	time.Sleep(time.Millisecond)
	print_num(n, "main")
}

func test_go2() {
	// 操作系统
	fmt.Println("操作系统:", runtime.GOOS)

	// cpu 数
	cpu_num := runtime.NumCPU()
	fmt.Println("系统cpu数:", cpu_num)

	// goroot目录
	go_root := runtime.GOROOT()
	fmt.Println("GOROOT目录:", go_root)

	// 让出时间片
	runtime.Gosched()
	fmt.Println("切换go")
}

func Pmain17() {
	go test_go()
	go test_go2()
	time.Sleep(time.Second)
}
