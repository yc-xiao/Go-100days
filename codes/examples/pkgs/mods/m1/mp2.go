package m1

import (
	"Go-100days/codes/mods/m2"
	"fmt"
)

func init() {
	fmt.Println("mp2")
	fmt.Println("add m1 ")
	m2.AddMap("m1", "m1_db")
}
