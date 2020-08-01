package main

import (
	"fmt"
	"golang.org/x/exp/mmap"
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

	mbUnit := 1024 * 1024 * 1024
	var block = make([]byte, mbUnit)
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

func testIoPkg(oldPath, newPath string) {
	lable := "testIoPkg"
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
	io.Copy(fWriter, fReader)
}

func testMMAPRW(oldPath, newPath string)  {
	lable := "testMMAPRW"

	fReader, err := mmap.Open(oldPath)
	if err != nil {
		log.Fatal(lable, " reader err ", err)
	}

	fWriter, err := os.OpenFile(newPath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Fatal(lable, " writer err ", err)
	}

	defer fReader.Close()
	defer fWriter.Close()

	mbUnit := 1024 * 1024 * 1024
	var block = make([]byte, mbUnit)
	var n int
	for {
		n, err = fReader.ReadAt(block, int64(n))
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

func main() {
	old, new := "/home/youcan/Desktop/1.tar.gz", "/home/youcan/Desktop/2.tar.gz"
	s0 := time.Now()
	testRW(old, new)
	s1 := time.Now()
	testIoPkg(old, new)
	s2 := time.Now()
	testMMAPRW(old, new)
	s3 := time.Now()

	fmt.Println("norm -> ", s1.Sub(s0))
	fmt.Println("iocp -> ", s2.Sub(s1))
	fmt.Println("mmap -> ", s3.Sub(s2))


}
