package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func testRW(oldPath, newPath string)  {
	lable := "testRW"
	fReader, err := os.Open(oldPath)
	if err != nil {
		log.Fatal(lable, " reader err ", err)
	}
	fWriter, err := os.OpenFile(newPath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Fatal(lable, " writer err ", err)
	}

	defer fReader.Close()
	defer fWriter.Close()
	//io.Copy(fReader, fWriter)
	//return

	var block = make([]byte, 1024 * 1024 * 1024)
	for {
		n, err := fReader.Read(block)
		if err != nil && err != io.EOF {
			log.Fatal(lable, " read err ", err)
		}
		if err == io.EOF || n == 0 {
			break
		}
		_, err = fWriter.Write(block[:n])
		if err != nil {
			log.Fatal(lable, " write err ", err)
		}
	}

}

func testMMAPRW(oldPath, newPath string)  {
	//lable := "testMMAPRW"
}

func main() {
	old, new := "/home/youcan/Desktop/tt.tar", "/home/youcan/Desktop/t2.tar"
	s := time.Now()
	testRW(old, new)
	s1 := time.Now()
	//testMMAPRW(old, new)
	s2 := time.Now()
	fmt.Println("mmap -> ", s2.Sub(s1), 's')
	fmt.Println("norm -> ", s1.Sub(s), 's')
}
