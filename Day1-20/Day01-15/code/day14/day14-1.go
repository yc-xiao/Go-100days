package day14

import "fmt"

type Persion struct {
	int
}

func (p Persion) add() {
	p.int++
}

func add(p Persion) {
	p.int++
}

func (p *Persion) add1() {
	p.int++
}

func add1(p *Persion) {
	p.int++
}

func Pmain14() {
	p := Persion{1}
	fmt.Println(p)
	p.add()
	fmt.Println(p)
	add(p)
	fmt.Println(p)

	p.add1()
	fmt.Println(p)
	add1(&p)
	fmt.Println(p)
}
