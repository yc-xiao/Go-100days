package main

import (
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strings"
)

func dumps(obj interface{}) ([]byte, error) {
	var s string
	objType := reflect.TypeOf(obj)
	objValue := reflect.ValueOf(obj)
	for i := 0; i < objType.NumField(); i++ {
		fieldType := objType.Field(i)
		objValue := objValue.Field(i)
		fmt.Println(fieldType.Name, fieldType.Type)
		fmt.Println(objValue)
		switch fieldType.Type.String() {
		case "string":
			s = s + fmt.Sprintf(`"%s": "%s",`, fieldType.Name, objValue)
		case "int":
			s = s + fmt.Sprintf(`"%s": %d,`, fieldType.Name, objValue)
		case "bool":
			s = s + fmt.Sprintf(`"%s": %t,`, fieldType.Name, objValue)
		default:
			return []byte{}, errors.New("dumps仅支持string/int/bool类型")
		}
	}
	s = "{" + strings.TrimRight(s, ",") + "}"
	return []byte(s), nil
}

func loads(obj interface{}, data []byte) error {
	return nil
}

func handlerErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	//　实现一个任意类型的json解析，并写到文件内
	type Persion struct {
		Name string
		Age  int
		Sex  bool
	}
	p := Persion{Name: "小米", Age: 20}
	data, err := dumps(p)
	handlerErr(err)
	ioutil.WriteFile("persion.json", data, os.ModePerm)
}
