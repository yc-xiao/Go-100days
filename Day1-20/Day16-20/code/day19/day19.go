package day19

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

type im interface {
	public()
}

type A struct {
	string
}

func (a A)public()  {
	fmt.Println("a.public")
}

func (a A)private(){
	fmt.Println(a.string)
}

type B struct {
	string
}

func (b B)public()  {
	fmt.Println("b.public")
}

func (b B)private(){
	fmt.Println(b.string)
}


func Pmain19() {
	test1()
	test2()
}


func test1()  {
	// stuct A，B 实现　interface i
	/*
		Go 所有变量都是静态类型， 而interface内部实现 (v, type) -> (nil, nil)
	*/

	var im1 im // interface -> (v, type) => (nil, nil)
	fmt.Printf("%v, %T, %p\n", im1, im1, &im1)

	im1 = A{"aaa"} // -> ({"aaa"}, A)
	fmt.Printf("%v, %T, %p\n", im1, im1, &im1)

	var im2 im = &A{"&aaa"} // ->(&A, A)
	fmt.Printf("%v, %T, %p\n", im2, im2, &im2)

	ia, ok := im1.(A)
	if ok {
		ia.private()
	}
	ib, ok := im1.(B)
	if ok {
		ib.private()
	}
	fmt.Println()
}

func test2()  {
	/*
		有时候不知道接口变量的　type 类型，可以使用反射操作
		断言也可以，但是type类型太多写断言麻烦!
	 */

	rand.Seed(int64(time.Now().Nanosecond()))
	num := rand.Intn(100)

	var im1 im

	if num >= 50{
		im1 = &A{"AAA"}
	}else {
		im1 = &B{"BBB"}
	}

	im1Type := reflect.TypeOf(im1) // 返回reflect -> type
	fmt.Println("reflect_type_obj:", im1Type)
	fmt.Println("kind():", im1Type.Kind())
	fmt.Println("Elem():", im1Type.Elem())
	fmt.Println()

	im1Value := reflect.ValueOf(im1) // 返回reflect -> value
	fmt.Println("reflect_value_obj:", im1Value)
	fmt.Println("Type():", im1Value.Type())
	fmt.Println("Kind():", im1Value.Kind())
	fmt.Println("Interface():", im1Value.Interface())
	fmt.Println("Elem():", im1Value.Elem())

}