package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func read() {
	f, err := os.Open("data.json")
	if err != nil {
		log.Fatal("open err -> ", err)
	}
	defer f.Close()
	data, err := ioutil.ReadAll(f)
	var p struct {
		// 大写
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	err = json.Unmarshal(data, &p)
	fmt.Println(err)
	fmt.Println(p)
	fmt.Println(string(data))
}

func write() {
	var p struct {
		Name  string `json:"name string"`
		Books []string
	}
	fmt.Println(p)
	p.Name = "小明"
	p.Books = []string{"booka", "bookb", "bookc"}
	data, err := json.Marshal(p)
	if err != nil {
		log.Fatal("json err", err)
	}
	fmt.Println(data)

	err = ioutil.WriteFile("data2.json", data, os.ModePerm)
	if err != nil {
		log.Fatal("open err", err)
	}
}

func main() {
	read()
	write()
}
