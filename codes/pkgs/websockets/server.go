package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func readHtml(path string) []byte {
	f, err := os.Open(path)
	if err != nil {
		fmt.Printf("open file(%s) err %v", path, err)
		return []byte{}
	}
	defer f.Close()

	buffer := bytes.NewBuffer([]byte{})
	buffer.ReadFrom(f)
	binary.Write(buffer, binary.LittleEndian, f)
	return buffer.Bytes()
}


func localHtml(resp http.ResponseWriter,  req *http.Request)  {
	data := readHtml("local.html")
	resp.Write(data)
}

func main() {
	router := mux.NewRouter()
	go h.run()
	router.HandleFunc("/ws", myws)
	router.HandleFunc("/local.html", localHtml)

	fmt.Println("监听 -> http://127.0.0.1:8080")
	fmt.Println("local.html -> http://127.0.0.1:8080/local.html")
	if err := http.ListenAndServe("127.0.0.1:8080", router); err != nil {
		fmt.Println("err:", err)
	}
}
