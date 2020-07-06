package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1,l2 := len(nums1), len(nums2)
	ll := l1+l2
	num3:=make([]int, ll)

	j,k:= 0,0
	for i:=0;i<ll;i++{
		if(j>=l1){
			num3[i] = nums2[k]
			k+=1
		}else if(k>=l2){
			num3[i] = nums1[j]
			j+=1
		}else if(nums1[j]>nums2[k]){
			num3[i] = nums2[k]
			k+=1
		}else {
			num3[i] = nums1[j]
			j+=1
		}
	}
	if ll%2 == 0{
		return float64(num3[ll/2-1] + num3[ll/2])/2
	}else {
		return float64(num3[ll/2])
	}
}

func main() {
	mid_num := findMedianSortedArrays([]int{1,2}, []int{3,4})
	fmt.Println(mid_num)
}