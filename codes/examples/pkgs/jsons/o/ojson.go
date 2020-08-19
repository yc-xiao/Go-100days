package o

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
)

func Loads(i interface{}, data []byte) {
	// 去掉 {}
	data = data[1:len(data)-1]
	datas := bytes.Split(data, []byte(","))
	for _, data = range datas{
		byteToValue(i, data)
	}
	// TODO
}

func byteToValue(i interface{}, data []byte) {

}

func Dumps(i interface{}) []byte{
	// 将 i 的值转换为二进制
	v := reflect.ValueOf(i)
	return []byte(getJsonString(v))
}

func getJsonString(v reflect.Value) string {
	switch v.Kind() {
	case reflect.String:
		return fmt.Sprintf(`"%s"`, v.String())
	case reflect.Bool:
		return fmt.Sprintf(`%t`, v.Bool())
	case reflect.Int:
		return fmt.Sprintf(`%d`, v.Int())
	case reflect.Int64:
		return fmt.Sprintf(`%d`, v.Int())
	case reflect.Float32:
		return fmt.Sprintf(`%f`, v.Float())
	case reflect.Float64:
		return fmt.Sprintf(`%f`, v.Float())
	case reflect.Array, reflect.Slice:
		ss := make([]string, v.Len())
		for i:=0; i < v.Len(); i++ {
			vv := v.Index(i)
			ss[i] = getJsonString(vv)
		}
		s := strings.Join(ss, ",")
		return fmt.Sprintf("[%s]", s)
	case reflect.Map:
		i := 0
		ss := make([]string, v.Len())
		iter := v.MapRange()
		for iter.Next() {
			vk, vv := iter.Key(), iter.Value()
			ss[i] = fmt.Sprintf(`"%s": %s`, vk.String(), getJsonString(vv))
			i++
		}
		s := strings.Join(ss, ",")
		return fmt.Sprintf("{%s}", s)
	case reflect.Interface:
		return getJsonString(v.Elem())
	case reflect.Ptr:
		return ""
	case reflect.Struct:
		l := v.NumField()
		ss := make([]string, l)
		for i:=0;i<l;i++{
			vattr := v.Type().Field(i)
			attr, ok := vattr.Tag.Lookup("json")
			if !ok {
				attr = vattr.Name
			}
			value := getJsonString(v.Field(i))
			ss[i] = fmt.Sprintf(`"%s": %s`, attr, value)
		}
		s := strings.Join(ss, ",")
		return fmt.Sprintf("{%s}", s)
	}
	return ""
}