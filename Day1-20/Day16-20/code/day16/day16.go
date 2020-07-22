package day16

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func create() {
	// 创建一个新的文件可读可写，若文件存在则清空
	os.Chdir("Day16-20/code/day16")
	path := "create.txt"
	file, err := os.Create(path)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	// 写入数据
	n, err := file.Write([]byte("正在创建文件!!!"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(n)

	// 读取数据， 需要调整文件偏移量
	// []byte{} 该切片大小未0, 读取数据需要指定切片大小 make([]byte, len, cap)
	data := make([]byte, n, n)
	n, err = file.ReadAt(data, 0)
	fmt.Println(n, data, string(data), err)

}

func open() {
	// 打开一个不存在的文件
	os.Chdir("Day16-20/code/day16")
	path := "open.txt"
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// 打开一个不存在文件，如果不存在则新增
	// os.O_TRUNC 清空文件
	// os.O_APPEND 追加文件
	file, err = os.OpenFile(path, os.O_RDWR|os.O_APPEND|os.O_TRUNC, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
	data := []byte("打开文件2333")
	file.Write(data)
}

func show() {
	dirname, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(dirname)

	file1, err := os.Stat(dirname)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(file1.IsDir())

	// 打开一个目录
	file2, err := os.Open(dirname)
	if err != nil {
		fmt.Println(err)
		return
	}
	fileInfos, err := file2.Readdir(100)
	if err != nil {
		fmt.Println(err)
	}
	for i, v := range fileInfos {
		fmt.Println(i, v)
	}
	//datas := make([]byte, 20 ,20)
	//n, err := file2.Read(datas)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(datas, n, err)
}

func util() {
	dirname, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}
	files, err := ioutil.ReadDir(dirname)
	for _, file := range files {
		fmt.Println(file.Name(), file.Size())
	}

}

func write() {
	os.Chdir("Day16-20/code/day16")
	dirname, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dirname)

	path := "text.txt"
	file, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_RDWR, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	wdatas := " <静夜思>\n   李白  \n床前明月光，\n疑是地上霜。\n举头望明月，\n低头思故乡。\n\n"
	wdatas = strings.Repeat(wdatas, 3)
	file.WriteString(wdatas)

	// 设置偏移量
	file.Seek(0, 0)
	rdatas := make([]byte, 100, 100)
	for {
		n, err := file.Read(rdatas)
		if n == 0 || err == io.EOF {
			break
		}
		fmt.Printf(string(rdatas[:n]))
	}
}

func Pmain16() {
	// 新建一个文件
	//create()
	// 打开一个文件
	//open()
	// 显示当前路径 以及当前目录下所有文件
	//show()
	// 文件工具包
	//util()
	// 写一个4M 数据
	//write()
}
