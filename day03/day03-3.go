package day03

import (
	"fmt"
	"math/rand"
	"time"
)

func Pmain33()  {
	rand.Seed(time.Now().UnixNano()) // 设置随机种子, 若不设置，每次运行结果都一致
	n := rand.Int()
	n1 := rand.Intn(100) // [0-100]
	n2 := rand.Intn(100) + 10 //[10-110]
	n3 := rand.Intn(100) // [0-100]
	fmt.Println(n, n1, n2, n3)

	t1 := time.Now()
	fmt.Println(t1)
	fmt.Println(t1.Day())
	fmt.Println(t1.Date())
	fmt.Println(t1.Unix())
	fmt.Println(t1.UnixNano())
	fmt.Printf("%T\n",t1)
}