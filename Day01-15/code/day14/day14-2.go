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
	fmt.Println("鼠标记录信息", m.info)
	m.info += info
	fmt.Println("鼠标记录信息", m.info)
	return m.info
}

func (m mouse) end() {
	fmt.Println("鼠标结束启动!")
	fmt.Println("鼠标记录信息", m.info)
}

func (d disk) start(info string) string {
	fmt.Println("硬盘开始启动!")
	fmt.Println("硬盘记录信息", d.info)
	d.info += info
	fmt.Println("硬盘记录信息", d.info)
	return d.info
}

func (d disk) end() {
	fmt.Println("硬盘结束启动!")
	fmt.Println("硬盘记录信息", d.info)
}

func Pmain14_2() {
	// 定义接口，并赋值
	var i process
	i = mouse{"鼠标信息!"}
	info := i.start("233333333")
	i.end()
	fmt.Println(info)

	// new
	var i2 process
	i2 = new(disk)
	info = i2.start("2333")
	i2.end()
	fmt.Println(info)
}
