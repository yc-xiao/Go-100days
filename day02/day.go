package day02

import (
	"fmt"
)

const NUMS = 1
const (
	a = iota   //0
	b          //1
	c          //2
	d = "ha"   //独立值，iota += 1
	e          //"ha"   iota += 1
	f = 100    //iota +=1
	g          //100  iota +=1
	h = iota   //7,恢复计数
	i          //8
)

// 打印常量
func Pnums()  {
	fmt.Println(NUMS)
	nums := [3]int{1,2,3}
	for i, v:=range nums{
		fmt.Printf("i=%d &i=%p, v=%d, &v=%p \n", i, &i, v, &v)
	}
	fmt.Println(a,b,c,d,e,f,g,h,i)
}