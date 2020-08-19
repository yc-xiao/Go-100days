package main

import (
	"fmt"
	"reflect"
)

var i interface{}

func testArrOrSclice() {
	arrs := []float64{1, 2, 3, 4, 5}
	i = arrs
	arrt := reflect.TypeOf(i)
	arrv := reflect.ValueOf(i)
	fmt.Println(arrt.Kind(), arrt.Elem().Kind())
	for i := 0; i < arrv.Len(); i++ {
		a := arrv.Index(i)
		fmt.Println(a.Kind(), a.Float())
	}
}

func testMap() {
	dic := map[string]float64{"1": 1, "2": 2, "3": 3}
	i = dic
	dict := reflect.TypeOf(i)
	dicv := reflect.ValueOf(i)
	fmt.Println(dict.Kind(), dict.Key().Kind(), dict.Elem().Kind())
	iter := dicv.MapRange()
	for iter.Next() {
		k, v := iter.Key(), iter.Value()
		fmt.Println(k.Kind(), k.String())
		fmt.Println(v.Kind(), v.Float())
	}
}

func testStruct() {
	type persion struct {
		Name string            `json:"name" label:"test"`
		arr  []string          `json:"arr" label:"test"`
		dic  map[string]string `json:"dic" label:"test"`
	}
	p := persion{"小米", []string{"1", "2", "3"}, map[string]string{"4": "4", "5": "5"}}
	i = p
	st := reflect.TypeOf(i)
	sv := reflect.ValueOf(i)
	fmt.Println(st.Kind())
	for i := 0; i < st.NumField(); i++ {
		t, v := st.Field(i), sv.Field(i)
		fmt.Println(t.Name, t.Type, t.Tag.Get("label"), v.Kind())
		switch v.Kind() {
		case reflect.String:
			fmt.Println(v.String())
		case reflect.Slice, reflect.Array:
			var s string
			for j := 0; j < v.Len(); j++ {
				vv := v.Index(j)
				s += vv.String()
			}
			fmt.Println(s)
		case reflect.Map:
			iter := v.MapRange()
			for iter.Next() {
				vk, vv := iter.Key(), iter.Value()
				fmt.Println(vk.Kind(), vk.String())
				fmt.Println(vv.Kind(), vv.String())
			}

		}
	}
}

func main() {
	testArrOrSclice()
	fmt.Println()
	testMap()
	fmt.Println()
	testStruct()
}
