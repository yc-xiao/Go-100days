package main

import "fmt"

func is_hw(s string) bool{
	ll := len(s)
	for i:=0;i<ll/2+1;i++{
		if s[i] != s[ll-i-1]{
			return false
		}
	}
	return true
}

func longestPalindrome(s string) string {
	// max_ll 回文长度
	// i, ll 左右点
	i, ll := 0, len(s)-1
	max_ll :=0
	hw := s
	// 两者距离少于最长回文数退出
	for ll-i+1>max_ll{

		right := ll
		// 左指针不动，移动右指针
		for right-i+1>max_ll{
			if s[i] == s[right]&&is_hw(s[i:right+1]) {
				if _max := right-i + 1; _max > max_ll {
					hw = s[i:right+1]
					max_ll = _max
					break
				}
			}
			right--
		}
		i+=1

		left:=i
		// 右指针不动，移动左指针
		for ll-left+1>max_ll{
			if s[left] == s[ll]&&is_hw(s[left:ll+1]) {
				if _max := ll-left + 1; _max > max_ll {
					hw = s[left:ll+1]
					max_ll = _max
					break
				}
			}
			left++
		}
		ll-=1
	}
	fmt.Println(hw)
	return hw
}

func main() {
	s:="aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabcaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	longestPalindrome(s)
}