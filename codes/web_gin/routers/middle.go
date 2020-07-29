package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func login(c *gin.Context) {
	if _, err := c.Cookie("pass"); err == nil {
		c.JSON(200, gin.H{"msg": "已有cookie"})
	} else {
		c.SetCookie("pass", "ok", 60, "/", "127.0.0.1", false, true)
		c.JSON(200, gin.H{"msg": "添加cookie"})
	}
}

func home(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "返回home"})
}

func logout(c *gin.Context) {
	c.SetCookie("pass", "ok", -1, "/", "127.0.0.1", false, false)
	c.JSON(200, gin.H{"msg": "取消cookie"})
}

func CheckCookies() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, err := c.Cookie("pass"); err == nil {
			c.Next() // 函数执行后在执行其他内容
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "err"})
			c.Abort() // 若验证不通过，不再调用后续的函数处理
		}
	}
}
