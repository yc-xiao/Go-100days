package day19

import (
	"fmt"
	"reflect"
)

type Persion struct {
	name string `a=2 b=4`
	age int `a=2 b=4`
	father *Persion
}

func (p Persion)who()  {
	fmt.Println("who is ", p.name)
}

func (p Persion)setAge(age int){
	p.age = age
}

func Pmain19_3()  {
	fmt.Println()
	// instance -> reflect.type, reflect.value
	p := Persion{name: "xm"}
	pt := reflect.TypeOf(p)
	ppt := reflect.TypeOf(&p)
	pv := reflect.ValueOf(p)
	ppv := reflect.ValueOf(&p)

	fmt.Printf("p -> %v &p -> %p T -> %T \n", p, &p, p)
	fmt.Printf("T: pt -> %T ppt -> %T \npv -> %T ppv -> %T \n\n", pt, ppt, pv, ppv)
	fmt.Printf("P: pt -> %p ppt -> %p \npv -> %p ppv -> %p \n\n", &pt, &ppt, &pv, &ppv)
	fmt.Printf("V: pt -> %v ppt -> %v \npv -> %v ppv -> %v \n\n", pt, ppt, pv, ppv)
	fmt.Println("Kind() -> p的类型 ", pt.Kind(), pv.Kind())
	fmt.Println("Kind() -> &p的类型 ", ppt.Kind(), ppv.Kind())
	fmt.Println()

	// reflect.type -> reflect.value ==> 生成一个新的空对象
	ptv, pptv := reflect.New(pt), reflect.New(ppt)
	fmt.Println(ptv, pptv.Kind())

	// reflect.value -> reflect.type ==> reflect.TypeOf
	pvt, ppvt := pv.Type(), ppv.Type()
	fmt.Printf("pvt -> T:%T  P:%p  v:%v\n", pvt, &pvt, pvt)
	fmt.Printf("ppvt -> T:%T  P:%p  v:%v\n\n", ppvt, &ppvt, ppvt)

	// reflect.value -> instance(interface) ==> reflect.TypeOf
	pvi, ppvi := pv.Interface(), ppv.Interface()
	pvi_obj, ppvi_obj := pvi.(Persion), ppvi.(*Persion)

	fmt.Printf("pvi -> T:%T  P:%p  v:%v\n", pvi, &pvi, pvi)
	fmt.Printf("ppvi -> T:%T  P:%p  v:%v\n\n", ppvi, &ppvi, ppvi)
	fmt.Printf("pvi_obj -> T:%T  P:%p  v:%v \n", pvi_obj, &pvi_obj, pvi_obj)
	fmt.Printf("ppvi_obj -> T:%T  P:%p  v:%v vp:%p\n\n", ppvi_obj, &ppvi_obj, ppvi_obj, ppvi_obj)

	// 指针才能使用Elem
	ppte, ppve := ppt.Elem(), ppv.Elem()
	fmt.Println(ppte, ppve)

	for i:=0;i<ppte.NumField();i++ {
		eft := ppte.Field(i)
		efv := ppve.Field(i)
		fmt.Println(eft.Name, eft.Type.Name())
		fmt.Println(eft, efv)
	}
	fmt.Println(ppv.CanSet(), ppve.CanSet())

}