package day15

import (
	"fmt"
	"os"
)

func Pmain15() {
	f, err := os.Open("hello.c")
	defer func() {
		fmt.Println(f)
		e := f.Close()
		fmt.Println(e)
	}()
	fmt.Println(f, err)
}
