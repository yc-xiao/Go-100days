package day13

import "fmt"

type A struct {
	name string
}

// B 继承A 拥有A的方法
type B struct {
	A
	name string
}

type C struct {
	name string
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

func Pmain13_2() {
	a := A{"a"}
	b := B{a, "b"}
	c := C{"c"}
	d := D{a, c}
	e := E{a}

	fmt.Println(a, b, c, d)
	// a a.name
	fmt.Println(a, a.name)
	// b -> a   b.name->b  b.A.name -> a
	fmt.Println(b, b.name, b.A.name)
	// c c.name
	fmt.Println(c, c.name)
	// d -> a,c  d.A.name -> a   d.C.name -> c
	fmt.Println(d, d.A.name, d.C.name)
	// e -> a   e.name -> a
	fmt.Println(e, e.name)

	e.publicFunc()
}
