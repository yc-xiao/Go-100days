package main

import "fmt"

func cmp(s1 string, s2 uint8) bool {
	for i := 0; i < len(s1); i++ {
		if s2 == s1[i] {
			return true
		}
	}
	return false
}

func lengthOfLongestSubstring(s string) int {
	ll := len(s)
	if ll < 2 {
		return ll
	}
	max := 0
	for i := 0; i < ll-1; i++ {
		var j int
		for j = i; j < ll-1; j++ {
			if cmp(s[i:j+1], s[j+1]) {
				break
			}
		}
		if _max := j + 1 - i; _max > max {
			max = _max
		}
		if ll-i-1 <= max {
			return max
		}
	}
	return max
}

func main() {
	map_string := map[string]int{"abcabcbb": 3, "bbbbb": 1, "pwwkew": 3, "aaa": 1, "": 0}
	for s, n := range map_string {
		if n1 := lengthOfLongestSubstring(s); n != n1 {
			fmt.Println(n, "!=", n1, s)
		} else {
			fmt.Println("ok")
		}
	}
}
