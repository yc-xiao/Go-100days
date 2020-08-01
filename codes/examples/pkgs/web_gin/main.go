package main

import (
	"Go-100days/codes/web_gin/routers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func f(c *gin.Context) {
	c.String(http.StatusOK, "hello World!")
}

func main() {
	r := gin.Default() // 1.创建路由
	r.GET("/", f)      // 2.绑定路由规则，执行的函数, gin.Context，封装了request和response
	routers.Register(r)
	r.Run() // 3.监听端口，默认在8080
}
