package main

import (
	"encoding/json"
	"fmt"
	"github.com/prometheus/common/log"
	"io/ioutil"
	"os"
)

type Persion struct {
	Name string `json:"name"`
	Age  int `json:"age"`
}

func handlerErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main()  {
	f, err := os.Open("data.json")
	handlerErr(err)
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	handlerErr(err)
	p := new(Persion)
	json.Unmarshal(data, p)
	fmt.Println(p)

	p.Name = "小明的弟弟"
	p.Age = p.Age - 2
	data, err = json.Marshal(p)
	handlerErr(err)

	ioutil.WriteFile("data2.json", data, os.ModePerm)
}

