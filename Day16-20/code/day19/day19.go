package day19

import (
	"fmt"
	"reflect"
)

func Pmain19() {
	n := 10
	ntype := reflect.TypeOf(n)
	fmt.Println(ntype)
	nvalue := reflect.ValueOf(n)
	fmt.Println(nvalue)
}
