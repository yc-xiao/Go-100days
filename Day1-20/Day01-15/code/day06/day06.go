package day06

import (
	"fmt"
	"math/rand"
	"time"
)

const NUM = 100

type private_sort func(arr [NUM]int) [NUM]int

// 冒泡排序
func mp_sort(arr [NUM]int) [NUM]int {
	for i := 0; i < NUM-1; i++ {
		for j := 0; j < NUM-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

// 插入排序
func insert_sort(arr [NUM]int) [NUM]int {
	// 从第一位开始向前排列
	for i := 1; i < NUM; i++ {
		n := i
		v := arr[n]
		// 当前位置向前比较
		for n > 0 {
			if v < arr[n-1] {
				arr[n] = arr[n-1]
			} else {
				break
			}
			n--
		}
		arr[n] = v
	}
	return arr
}

// 选择排序
func select_sort(arr [NUM]int) [NUM]int {
	for i := 0; i < NUM-1; i++ {
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

func _quick_sort(arr *[NUM]int, left, right int) {
	// 结束条件
	if left+1 >= right {
		if arr[left] > arr[right] {
			arr[right], arr[left] = arr[left], arr[right]
		}
		return
	}

	// 从左到右，从右到左，不相遇
	lindex, rindex := left+1, right
	for lindex < rindex {
		for lindex < rindex {
			if arr[lindex] <= arr[left] {
				lindex++
			} else {
				break
			}
		}
		for lindex < rindex {
			if arr[rindex] >= arr[left] {
				rindex--
			} else {
				break
			}
		}
		if lindex < rindex {
			arr[lindex], arr[rindex] = arr[rindex], arr[lindex]
		}
	}

	if arr[left] > arr[lindex] {
		arr[lindex], arr[left] = arr[left], arr[lindex]
	}
	_quick_sort(arr, left, lindex-1)
	_quick_sort(arr, rindex, right)
}

// 快速排序
func quick_sort(arr [NUM]int) [NUM]int {
	// 传数组，闭区间
	_quick_sort(&arr, 0, NUM-1)
	return arr
}

func _two_sort(arr []int) []int {
	// 递归返回条件
	ll := len(arr)
	if ll <= 2 {
		if ll == 2 && arr[1] < arr[0] {
			return []int{arr[1], arr[0]}
		}
		return arr
	}
	// 拆分
	mid := ll / 2
	left_arr := _two_sort(arr[:mid])
	right_arr := _two_sort(arr[mid:ll])

	lindex, rindex, lend, rend := 0, 0, len(left_arr), len(right_arr)

	// 合并，切片多次生成新的空间不利于
	//end_arr := []int{}
	//for i := 0; i < ll; i++ {
	//	switch  {
	//	case lindex >= lend:
	//		end_arr = append(end_arr, right_arr[rindex])
	//		rindex++
	//	case rindex >= rend:
	//		end_arr = append(end_arr, left_arr[lindex])
	//		lindex++
	//	case left_arr[lindex] > right_arr[rindex]:
	//		end_arr = append(end_arr, right_arr[rindex])
	//		rindex++
	//	case left_arr[lindex] <= right_arr[rindex]:
	//		end_arr = append(end_arr, left_arr[lindex])
	//		lindex++
	//	}
	//}

	// 生成指定长度的切片
	end_arr := make([]int, ll)
	for i := 0; i < ll; i++ {
		if lindex >= lend {
			end_arr[i] = right_arr[rindex]
			rindex++
		} else if rindex >= rend {
			end_arr[i] = left_arr[lindex]
			lindex++
		} else if left_arr[lindex] > right_arr[rindex] {
			end_arr[i] = right_arr[rindex]
			rindex++
		} else if left_arr[lindex] <= right_arr[rindex] {
			end_arr[i] = left_arr[lindex]
			lindex++
		}
	}
	return end_arr
}

// 归并排序
func two_sort(arr [NUM]int) [NUM]int {
	slice_arr := _two_sort(arr[:])
	for i := 0; i < NUM; i++ {
		arr[i] = slice_arr[i]
	}
	return arr
}

// 生成随机数
func generate(arr *[NUM]int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < NUM; i++ {
		arr[i] = rand.Intn(NUM)
	}
	fmt.Printf("随机生成数组arr:%v\n\n", arr)
}

// 打印函数执行时间
func pfuntime(sort private_sort, arr [NUM]int, sort_name string) [NUM]int {
	t1 := time.Now()
	out_arr := sort(arr)
	t2 := time.Now()
	fmt.Println(sort_name, "执行时间:", t2.Sub(t1))

	if NUM <= 500 {
		fmt.Printf("初始参数: %v\n", arr)
		fmt.Printf("执行结果: %v\n\n", out_arr)
	}
	return out_arr
}

func Pmain6() {
	var arr [NUM]int
	generate(&arr)

	_ = pfuntime(mp_sort, arr, "冒泡排序")     // 冒泡排序
	_ = pfuntime(insert_sort, arr, "插入排序") // 插入排序
	_ = pfuntime(select_sort, arr, "选择排序") // 选择排序
	_ = pfuntime(quick_sort, arr, "快速排序")  // 快速排序
	_ = pfuntime(two_sort, arr, "归并排序")    // 归并排序
	//heap_arr := heap_sort(arr) // 堆排序
}
