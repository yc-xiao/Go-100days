package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getRandNums(start, end int) []int {
	rand.Seed(time.Now().UnixNano())
	ll := end - start
	nums := make([]int, ll)
	for i:=0; i< ll; i++ {
		nums[i] = rand.Intn(ll) + start
	}
	return nums
}

type sort func(arr []int) []int

func Sort(f sort, arr []int, text string) {
	newArr := make([]int, len(arr))
	copy(newArr, arr)
	start := time.Now()
	newArr = f(newArr)
	finish := time.Now()
	fmt.Printf("%s，共花了%v秒\n", text, finish.Sub(start).Seconds())
	if len(arr) <= 1000 {
		fmt.Println(arr)
		fmt.Println(newArr)
	}
}

func bubbleSort(arr []int) []int{
	ll := len(arr)
	for j:=0; j < ll-1; j++{
		for i:=0; i < ll-j-1; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
	}
	return arr
}

func selectSort(arr []int) []int{
	ll := len(arr)
	for i:=0; i < ll-1; i++ {
		index := i
		for j:=i+1; j <= ll-1; j++ {
			if arr[index] > arr[j] {
				index, j = j, index
			}
		}
		arr[index], arr[i] = arr[i], arr[index]
	}
	return arr
}

func insertSort(arr []int) []int{
	ll := len(arr)
	for i:=1; i<ll; i++ {
		j := i
		for j>0 && arr[j] <= arr[j-1]{
			arr[j], arr[j-1] = arr[j-1], arr[j]
			j--
		}
	}
	return arr
}

func mergeSort(arr []int)  []int{
	/*
		归并排序：
			1. 将数组拆成两部分，每一部分进行递归
			2. 将两部分合并
	*/
	ll := len(arr)
	if ll < 2 { return arr}

	mid := ll/2
	leftArr := arr[0:mid]
	rightArr := arr[mid:ll]
	leftArr = mergeSort(leftArr)
	rightArr = mergeSort(rightArr)

	newArr :=  make([]int, ll)
	lindex,  rindex, lmax, rmax := 0, 0, mid, ll-mid
	for i:=0; i < ll; i++ {
		switch  {
		case lindex >= lmax:
			newArr[i] = rightArr[rindex]
			rindex++
		case rindex >= rmax:
			newArr[i] = leftArr[lindex]
			lindex++
		case leftArr[lindex] > rightArr[rindex]:
			newArr[i] = rightArr[rindex]
			rindex++
		case leftArr[lindex] <= rightArr[rindex]:
			newArr[i] = leftArr[lindex]
			lindex++
		}
	}
	return newArr
}


func main() {
	arr := getRandNums(0, 10)
	Sort(bubbleSort, arr, "冒泡排序")
	Sort(selectSort, arr, "选择排序")
	Sort(insertSort, arr, "插入排序")
	Sort(mergeSort, arr, "归并排序")
}

