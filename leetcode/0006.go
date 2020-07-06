package main
import "fmt"

func split(s string, n int) (string, string){
	n+=1
	new_s, split_s:="", ""
	for i:=0;i<len(s);i++{
		if i%n == 0{
			new_s+=string(s[i])
		}else {
			split_s+=string(s[i])
		}
	}
	return new_s, split_s
}

func convert(s string, numRows int) string {
	if numRows==1{
		return s
	}
	step := 2*(numRows-2)+1
	new_s, split_s := split(s, step)
	return new_s + convert(split_s, numRows-1)
}

func main() {

	// 输入: s = "LEETCODEISHIRING", numRows = 3
	//输出: "LCIRETOESIIGEDHN"
	// 输入: s = "LEETCODEISHIRING", numRows = 4
	//输出: "LDREOEIIECIHNTSG"
	//result:=convert("LEETCODEISHIRING", 3)
	//fmt.Println(result)
	s1:=convert("LEETCODEISHIRING", 3)
	if s1 =="LCIRETOESIIGEDHN"{
		fmt.Println("ok")
	}
	fmt.Println(s1)
}