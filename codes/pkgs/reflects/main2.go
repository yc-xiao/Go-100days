package main

import (
	"fmt"
	"reflect"
)

type Persion struct {
	Name string
	Age int
	Son *Persion
}

func (p *Persion)addAge() {
	p.Age++
	if p.Son != nil {
		p.Son.Age++
	}
}

func (p *Persion)String() string{
	return fmt.Sprintf("name -> %s, age -> %d, son -> %p", p.Name, p.Age, p.Son)
}

func pp(p *Persion){
	fmt.Println(p)
	if p.Son != nil{
		pp(p.Son)
	}
}

func run(obj interface{}){
	objType := reflect.TypeOf(obj)
	objValue := reflect.ValueOf(obj)
	if objType.Kind() != reflect.Ptr{
		fmt.Println("obj -> ", objType.Kind(), objValue)
	}
	fmt.Println("obj.Kind()->", objType.Kind())
	objt := objType.Elem()
	objv := objValue.Elem()
	fmt.Println(objv.CanSet())
	for i:=0; i < objt.NumField(); i++ {
		t := objt.Field(i)
		v := objv.Field(i)
		switch t.Name {
		case "Name":
			v.SetString("小明2")
		case "Age":
			v.SetInt(23)
		}
	}

}

func main() {
	p1 := &Persion{"老明", 50, &Persion{"小明", 20, nil}}
	p1.addAge()
	pp(p1)
	run(p1)
	pp(p1)
}


