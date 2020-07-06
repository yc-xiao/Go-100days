package day06

import (
	"fmt"
	"math/rand"
	"time"
)

const NUM = 100
type private_sort func(arr [NUM]int)([NUM]int)

// 冒泡排序
func mp_sort(arr [NUM]int) ([NUM]int){
	for i:=0;i<NUM-1;i++{
		for j:=0;j<NUM-1-i;j++{
			if arr[j]>arr[j+1]{
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

// 插入排序
func insert_sort(arr [NUM]int) ([NUM]int){
	return arr
}

// 选择排序
func select_sort(arr [NUM]int) ([NUM]int){
	for i := 0; i < NUM-1 ; i++ {
		max := 0
		for j := 0; j < NUM-i; j++ {
			if arr[j] > arr[max] {
				max = j
			}
		}
		arr[max], arr[NUM-i-1] = arr[NUM-i-1], arr[max]
	}
	return arr
}

// 快速排序
func quick_sort(arr [NUM]int) ([NUM]int){
	return arr
}

// 归并排序
func two_sort(arr [NUM]int) ([NUM]int){
	return arr
}

// 生成随机数
func generate(arr *[NUM]int){
	rand.Seed(time.Now().UnixNano())
	for i:=0;i<NUM;i++{
		arr[i] = rand.Intn(NUM)
	}
	fmt.Printf("随机生成数组arr:%v\n\n", arr)
}

// 打印函数执行时间
func pfuntime(sort private_sort, arr [NUM]int, sort_name string)([NUM]int){
	t1 := time.Now()
	out_arr:=sort(arr)
	t2 := time.Now()
	fmt.Println(sort_name, "执行时间:", t2.Sub(t1))
	fmt.Printf("初始参数: %v\n", arr)
	fmt.Printf("执行结果: %v\n\n", out_arr)
	return out_arr
}

func Pmain6(){
	var arr [NUM]int
	generate(&arr)

	_ = pfuntime(mp_sort, arr, "冒泡排序") // 冒泡排序
	_ = pfuntime(insert_sort, arr, "插入排序") // 插入排序
	_ = pfuntime(select_sort, arr, "选择排序") // 选择排序
	_ = pfuntime(quick_sort, arr, "快速排序") // 快速排序
	_ = pfuntime(two_sort, arr, "归并排序") // 归并排序
	//heap_arr := heap_sort(arr) // 堆排序
}