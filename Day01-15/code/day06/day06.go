package day06

import (
	"fmt"
	"math/rand"
	"time"
)

const NUMS = 100

// 冒泡排序
func mp_sort(arr [NUMS]int) ([NUMS]int){
	fmt.Println(arr)
	return arr
}

// 生成随机数
func generate(arr *[NUMS]int){
	rand.Seed(time.Now().UnixNano())
	for i:=0;i<NUMS;i++{
		arr[i] = rand.Intn(NUMS)
	}
}

// 打印函数执行时间
func pfuntime(){
	t1 := time.Now().Unix()
	t2 := time.Now().Unix()
	fmt.Println("函数执行时间:", t2-t1)
}

func Pmain6(){
	var arr [NUMS]int
	generate(&arr)

	mp_arr := mp_sort(arr) // 冒泡排序
	//insert_arr := insert_sort(arr) // 插入排序
	//select_arr := select_sort(arr) // 选择排序
	//quick_arr := quick_sort(arr) // 快速排序
	//heap_arr := heap_sort(arr) // 堆排序
	//two_arr := two_sort(arr) // 归并排序
	fmt.Println(mp_arr)
}