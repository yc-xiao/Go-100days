package examples

import "fmt"

func Func5() {
	/*
		== 相同类型比较
	*/
	sn1 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}
	sn2 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}

	if sn1 == sn2 {
		fmt.Println("sn1 == sn2")
	}

	sm1 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}
	sm2 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}
	fmt.Println(sm1, sm2)
	//if sm1 == sm2 {
	//	fmt.Println("sm1 == sm2")
	//}
}
