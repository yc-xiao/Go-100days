package main

import "fmt"

func intToRoman(num int) string {
	map_dic := map[int]string{1: "I", 4: "IV", 5: "V", 9: "IX", 10: "X", 40: "XL",
		50: "L", 90: "XC", 100: "C", 400: "CD", 500: "D", 900: "CM", 1000: "M"}
	str := ""
	n10 := 1
	for num > 0 {
		end_num := num % 10 * n10
		v, ok := map_dic[end_num]
		if !ok {
			v = ""
			for end_num > 0 {
				if end_num > 5*n10 {
					v += map_dic[5*n10]
					end_num -= 5 * n10
				} else {
					v += map_dic[n10]
					end_num -= n10
				}
			}
		}
		str = v + str
		n10 *= 10
		num /= 10
	}
	return str
}

func main() {
	dic := map[int]string{
		3:    "III",
		4:    "IV",
		9:    "IX",
		58:   "LVIII",
		1994: "MCMXCIV",
	}
	for n, s := range dic {
		if ss := intToRoman(n); ss != s {
			fmt.Println("error", n, s, ss)
		}
	}
}
