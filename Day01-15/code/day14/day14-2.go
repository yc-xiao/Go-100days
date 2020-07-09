package day14

import "fmt"

// 定义接口
type process interface {
	start(string) string
	end()
}

type mouse struct {
	info string
}

type disk struct {
	info string
}

func (m mouse) start(info string) string {
	fmt.Println("鼠标开始启动!")
	m.info += info
	return m.info
}

func (m mouse) end() {
	fmt.Println("鼠标结束启动!")
	fmt.Println("鼠标记录信息", m.info)
}

func (d disk) start(info string) string {
	fmt.Println("硬盘开始启动!")
	d.info += info
	return d.info
}

func (d disk) end() {
	fmt.Println("硬盘结束启动!")
	fmt.Println("硬盘记录信息", d.info)
}

func test_inferface1() {
	// 定义接口，并赋值
	var i process
	i = mouse{"鼠标信息!"}
	info := i.start("233333333")
	i.end()
	fmt.Printf("%T, ", i)

	// new
	var i2 process
	i2 = new(disk)
	info = i2.start("2333")
	i2.end()
	fmt.Println(info)
}

func test_inferface2() {
	test_inferface(1)
	test_inferface("2")
	test_inferface([]int{1, 2, 3})
}

func test_inferface(i interface{}) {
	fmt.Printf("%T, %v \n", i, i)
}

func test_inferface3() {
	dic := make(map[int]interface{})
	dic[0] = 1
	dic[1] = true
	dic[2] = "aaaa"
	dic[3] = mouse{"helloc"}
	// 无法调用值或对应类型所属方法
	// int(dic[0])
	fmt.Println(dic)
	fmt.Printf("dic[1], %T, \n", dic[1])
	fmt.Printf("dic[2], %T, \n", dic[2])
}

func Pmain14_2() {
	test_inferface1()
	test_inferface2()
	test_inferface3()
}
