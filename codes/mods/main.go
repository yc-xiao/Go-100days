package main

import (
	_ "Go-100days/codes/mods/m1"
	"Go-100days/codes/mods/m2"
	"fmt"
)

func init() {
	fmt.Println("import m1 ==> m2.Register(m1.key, m1.value) ==> tempMap")
	fmt.Println("import m2 ==> m2.GetMap(m1.key) ==> m1.value ==> tempMap")
}

func main() {
	// 导包逻辑
	value, ok := m2.GetMap("m1")
	fmt.Println(value, ok)
}