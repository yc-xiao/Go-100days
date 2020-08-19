package o

import (
	"bytes"
	"fmt"
	"reflect"
)

func dumps(t reflect.Type, v reflect.Value) []byte{
	buf := bytes.NewBuffer([]byte{})
	var s string
	switch t.Kind() {
	case reflect.String:
		s = fmt.Sprintf(`"%s"`, v.String())
		buf.WriteString(s)
	case reflect.Int:
		s = fmt.Sprintf(`%d`, v.Int())
		buf.WriteString(s)
	case reflect.Map:
		buf.WriteString("{")
		i, l := 1, v.Len()
		iter := v.MapRange()
		for iter.Next() {
			s = fmt.Sprintf(`"%s":`, iter.Key())
			buf.WriteString(s)
			value := Dumps(iter.Value().Interface())
			buf.Write(value)
			if i != l {
				buf.WriteString(",")
			}
			i++
		}
		buf.WriteString("}")
	case reflect.Array, reflect.Slice:
		s = fmt.Sprintf("%v", v)
		buf.WriteString(s)
	case reflect.Struct:
		buf.WriteString("{")
		l := t.NumField() - 1
		for i := 0; i < t.NumField(); i++ {
			Tfield := t.Field(i)
			Vfield := v.Field(i)
			tag := Tfield.Tag.Get("json")
			changeValue := dumps(Vfield.Type(), Vfield)
			if l == i {
				s = fmt.Sprintf(`"%s": %s`, tag, string(changeValue))
			}else{
				s = fmt.Sprintf(`"%s": %s,`, tag, string(changeValue))
			}
			buf.WriteString(s)
		}
		buf.WriteString("}")
	}
	return buf.Bytes()
}

func Dumps(i interface{}) []byte{
	// 将 i 的值转换为二进制
	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)
	return dumps(t, v)
}

//func Loads(i interface{}, data []byte) {
//	// 将二进制`{"a":1}` 赋予　i
//}
