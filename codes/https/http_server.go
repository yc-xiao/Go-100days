package https

import (
	"fmt"
	"net/http"
)

func HttpServer() {
	addr := "localhost:9000"
	fmt.Println("启动http_server, 监听 -> ", addr)
	http.HandleFunc("/", myHandler)
	http.ListenAndServe("localhost:9000", nil)
}

func myHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Println("url -> ", request.URL)
	fmt.Println("remoteAddr -> ", request.RemoteAddr)
	fmt.Println("header -> ", request.Header)
	fmt.Println("method -> ", request.Method)
	data := make([]byte, 1024)
	n, err := request.Body.Read(data)
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println("body -> ", data[:n])
	response.Write([]byte("小伙子你又来啦!!!"))
}
