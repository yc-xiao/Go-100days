package day13

import "fmt"

type A struct {
	int
}

// B 继承A 拥有A的方法
type B struct {
	A
	int
}

type C struct {
	int
}

// D 继承　A,C
type D struct {
	A
	C
}

// E 继承 A
type E struct {
	A
}

/*
	A <- B
	(A,C) <- D
*/

func (a A) publicFunc() {
	fmt.Printf("%v publicFunc\n", a)
}

func Pmain131() {
	a := A{10}
	b := B{a, 20}
	c := C{30}
	d := D{a, c}
	e := E{a}

	fmt.Println(a, b, c, d)
	fmt.Println(a, a.int)
	fmt.Println(b, b.int, b.A.int)
	fmt.Println(c, c.int)
	fmt.Println(d, d.A.int, d.C.int)
	fmt.Println(e, e.int)
}
