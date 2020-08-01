package main

import (
	"fmt"
	"testing"
)

type p struct {
	s string
	n int
}

func (o *p) print() {
	fmt.Println(o.s, o.n)
}

func TestStructInit(t *testing.T) {
	var p1 p
	var p2 = p{s: "p2"}
	p1.n = 1
	p2.n = 2
	p1.print()
	p2.print()
}
