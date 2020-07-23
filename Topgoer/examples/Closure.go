package examples

import "fmt"

func Func20C() {
	type Person struct {
		age int
	}
	person := &Person{28}

	// 1.
	defer fmt.Println(3, person.age) // -> 28

	// 2.
	defer func(p *Person) {
		fmt.Println(2, p.age) // -> 28
	}(person)

	// 3.
	defer func() {
		fmt.Println(1, person.age) // -> 29
	}()

	person = &Person{29}
}

func Func19C() {
	// 1. 先复制值，保存
	// 2. 先在本地找，再去外面找
	// 3. 是否是指针引用
	type Person struct {
		age int
	}
	person := &Person{28}

	defer fmt.Println(3, person.age) //  复制值到本地 -> 28

	defer func(p *Person) {
		fmt.Println(2, p.age) // 现在本地找，指针应用 -> 29
	}(person)

	defer func() {
		fmt.Println(1, person.age) // 先在本地找，再去外面找　-> 29
	}()

	person.age = 29
}

func Func18C() {
	f1 := func() (r int) {
		defer func() {
			r++
		}()
		return 0
	}

	f2 := func() (r int) {
		t := 5
		defer func() {
			t = t + 5
		}()
		return t
	}

	f3 := func() (r int) {
		defer func(r int) {
			r = r + 5
		}(r)
		return 1
	}
	fmt.Println(f1()) // 1  r=0, r=r+1, return r
	fmt.Println(f2()) // 5  r=5, return r
	fmt.Println(f3()) // 1  r=0, return r
}

func Func17C() {
	increaseA := func() int {
		var i int
		defer func() {
			i++
			fmt.Printf("i ")
		}()
		return i
	}

	increaseB := func() (r int) {
		defer func() {
			r++
			fmt.Printf("r ")
		}()
		return 1
	}
	fmt.Println(increaseA()) // 0 => 返回0
	fmt.Println(increaseB()) // 2 => 返回r
}
