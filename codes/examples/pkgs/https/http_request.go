package https

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
)

func HttpRequest(url string) {
	lable := "HttpRequest"
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(lable, err)
	}
	defer res.Body.Close()

	fmt.Println(lable, "Header Content-Type -> ", res.Header["Content-Type"])
	fmt.Println(lable, "Method -> ", res.Request.Method)
	fmt.Println(lable, "status -> ", res.Status)
	buf := make([]byte, 1024, 1024)
	bufs := bytes.NewBuffer([]byte{})

	for {
		n, err := res.Body.Read(buf)
		bufs.Write(buf[:n])
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}
	}
	fmt.Println(lable, "body len -> ", bufs.Len())
	//fmt.Println(string(bufs.Bytes()))
}

func TestHttpRequest() {
	HttpRequest("http://www.baidu.com")
}
