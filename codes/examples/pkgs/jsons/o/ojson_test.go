package o

import (
	"fmt"
	"testing"
)

func TestDumps(t *testing.T) {
	type persion struct {
		Name string `json:"name"`
		Age int `json:"age"`
		arr []int `json:"arr"`
	}
	h := map[string]interface{}{}
	h["hello"] = "c"
	h["age"] = 10
	h["arr"] = [...]int{1,2,3}
	h["p"] = persion{"小米", 10, []int{1,2,3}}
	data := Dumps(h)
	fmt.Println(string(data))
}
