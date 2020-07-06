package main

import (
	"fmt"
	"strconv"
)

//给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复
//的三元组。
// 示例：
// 给定数组 nums = [-1, 0, 1, 2, -1, -4]，
//满足要求的三元组集合为：
//[
//  [-1, 0, 1],
//  [-1, -1, 2]
func threeSum(nums []int) [][]int {
	nums = sort_set(nums)
	arrs := [][]int{}
	arr_set := make(map[string]int)

	ll := len(nums)
	var temp_arr []int
	for i:=0;i<ll-2;i++ {
		for j := i + 1; j < ll-1; j++ {

			n1 := strconv.Itoa(nums[i])
			n2 := strconv.Itoa(nums[j])
			_, ok := arr_set[n1+"-"+n2]
			if ok {
				continue
			}

			for k := j + 1; k < ll; k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					//n3 := strconv.Itoa(nums[k])
					temp_arr = []int{nums[i], nums[j], nums[k]}
					arrs = append(arrs, temp_arr)

					arr_set[n1+"-"+n2] = nums[k]
					//arr_set[n2+"-"+n1] = nums[k]
					//arr_set[n1+"-"+n3] = nums[j]
					//arr_set[n3+"-"+n1] = nums[j]
					//arr_set[n2+"-"+n3] = nums[i]
					//arr_set[n3+"-"+n2] = nums[i]

					_, ok := arr_set[n1+"-"+n2]
					if ok {
						break
					}
				}
			}
		}
	}
	return arrs
}


func sort_set(nums[]int)  []int{
	ll := len(nums)
	for i:=0;i<ll;i++{
		for j:=i+1;j<ll;j++{
			if  nums[i]>nums[j]{
				nums[i],nums[j] = nums[j],nums[i]
			}
		}
	}
	arr := []int{}
	dic := map[int]int{}
	for i:=0;i<ll;i++{
		_, ok := dic[nums[i]]
		if ok {
			continue
		}
		dic[nums[i]] = i
		arr = append(arr, nums[i])
	}
	return arr
}

func main() {
	nums := []int{9,-15,6,6,10,-2,8,8,0,-6,-4,-2,14,-6,-14,-2,12,5,-2,-8,13,13,-10,4,-6,8,6,-15,-5,11,-15,11,3,-2,-6,-10,11,-12,13,-12,-11,-5,2,10,-4,-5,-15,-7,7,-2,0,5,-11,-3,-13,-10,-9,0,-10,-7,-12,12,-11,7,-5,-1,12,-8,-6,3,-13,-10,5,-4,-14,-14,6,8,-14,-9,-8,-7,3,-4,6,5,1,12,-9,6,-10,-6,-5,-14,-14,5,-8,6,4,1}
	println(len(nums))
	nums = sort_set(nums)
	println(len(nums))
	fmt.Println(nums)
	result:=threeSum(nums)
	for i, v := range result{
		fmt.Println(v, i)
	}
	fmt.Println(len(result))
}