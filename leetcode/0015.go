package main

import (
	"fmt"
)

func threeSum(nums []int) [][]int {
	nums = quick_sort(nums, 0, len(nums)-1)
	arrs := [][]int{}
	arr_set := make(map[int]int)

	ll := len(nums)
	var temp_arr []int
	for i := 0; i < ll-2; i++ {
		for j := i + 1; j < ll-1; j++ {
			key := nums[i]*1000 + nums[j]
			if _, ok := arr_set[key]; ok {
				continue
			}
			for k := j + 1; k < ll; k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					temp_arr = []int{nums[i], nums[j], nums[k]}
					arrs = append(arrs, temp_arr)
					arr_set[key] = nums[k]
					break
				}
			}
		}
	}
	return arrs
}

func quick_sort(nums []int, left, right int) []int {
	// nums[left, right] 闭区间
	// 退出条件
	if right <= left {
		return nums
	}
	if right-left <= 1 {
		if nums[left] > nums[right] {
			nums[left], nums[right] = nums[right], nums[left]
		}
		return nums
	}

	// 进行排序
	lindex, rindex := left+1, right
	for lindex < rindex {
		// 左移动
		for lindex < rindex {
			if nums[left] >= nums[lindex] {
				lindex++
			} else {
				break
			}
		}
		// 右移动
		for lindex < rindex {
			if nums[left] <= nums[rindex] {
				rindex--
			} else {
				break
			}
		}

		if lindex >= rindex {
			break
		} else {
			nums[lindex], nums[rindex] = nums[rindex], nums[lindex]
		}
	}

	if nums[left] > nums[lindex] {
		nums[left], nums[lindex] = nums[lindex], nums[left]
	}

	quick_sort(nums, left, lindex-1) // 左侧排序
	quick_sort(nums, rindex, right)  // 右侧排序
	return nums
}

func main() {
	// 超时!!!
	nums := []int{9, -15, 6, 6, 10, -2, 8, 8, 0, -6, -4, -2, 14, -6, -14, -2, 12, 5, -2, -8, 13, 13, -10, 4, -6, 8, 6, -15, -5, 11, -15, 11, 3, -2, -6, -10, 11, -12, 13, -12, -11, -5, 2, 10, -4, -5, -15, -7, 7, -2, 0, 5, -11, -3, -13, -10, -9, 0, -10, -7, -12, 12, -11, 7, -5, -1, 12, -8, -6, 3, -13, -10, 5, -4, -14, -14, 6, 8, -14, -9, -8, -7, 3, -4, 6, 5, 1, 12, -9, 6, -10, -6, -5, -14, -14, 5, -8, 6, 4, 1}
	result := threeSum(nums)
	fmt.Println(len(result))
	fmt.Println(result)
}
