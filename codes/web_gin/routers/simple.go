package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func urlParam(c *gin.Context) {
	// /test1/小明/正/在准备跑
	name := c.Param("name")
	action := c.Param("action")
	action = strings.Trim(action, "/")
	c.String(http.StatusOK, name+" is "+action)
}

func getQuery(c *gin.Context) {
	// /test2/ or /test2/?name="小黄"
	name := c.DefaultQuery("name", "小明")
	c.String(http.StatusOK, fmt.Sprintf("hello %s", name))
}

func postForm(c *gin.Context) {
	types := c.DefaultPostForm("type", "post")
	username := c.PostForm("username")
	password := c.PostForm("password")
	c.String(http.StatusOK, fmt.Sprintf("username:%s,password:%s,type:%s", username, password, types))
}

func uploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.String(500, "上传图片出错")
	}
	// c.JSON(200, gin.H{"message": file.Header.Context})
	c.SaveUploadedFile(file, file.Filename)
	c.String(http.StatusOK, file.Filename)
}
