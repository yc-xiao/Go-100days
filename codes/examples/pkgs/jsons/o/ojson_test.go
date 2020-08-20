package o

import (
	"fmt"
	"testing"
)

func TestDumps(t *testing.T) {
	type persion struct {
		Name string `json:"name"`
		Age  int
		Son  []string `json:"son"`
	}
	h := map[string]interface{}{
		"int":   1,
		"array": [...]int{1, 2, 3},
		"dict": map[string]string{
			"a": "a1",
			"b": "b1",
		},
		"struct": persion{
			Name: "小明",
			Age:  18,
			Son:  []string{"小小明", "小小小明"},
		},
		"prt": &persion{
			Name: "小红",
			Age:  18,
			Son:  []string{"小小红", "小小小红"},
		},
	}

	data := Dumps(h)
	fmt.Println(string(data))
	fmt.Printf("%v", data)
}

func TestLoads(t *testing.T) {
	type persion struct {
		Name string   `json:"name"`
		Age  int      `json:"age"`
		Son  []string `json:"son"`
	}
	type Obj struct {
		Int     int               `json:"int"`
		array   []int             `json:"array"`
		dict    map[string]string `json:"dict"`
		persion `json:"persion"`
	}
	data := []byte(`{"int": 1,"array": [1,2,3],"dict": {"a": "a1","b": "b1"},"struct": {"name": "小明","age": 18,"son": ["小小明","小小小明"]}}
`)

	obj := new(Obj)
	Loads(obj, data)
	fmt.Println(obj)
}
