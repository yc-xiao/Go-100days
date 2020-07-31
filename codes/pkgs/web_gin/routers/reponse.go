package routers

import "github.com/gin-gonic/gin"

func returnText(c *gin.Context) {
	c.String(200, "返回一条信息")
}

func returnJson(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "返回一个json"})
}

func returnXml(c *gin.Context) {
	c.XML(200, gin.H{"msg": "返回一个xml"})
}

func returnYaml(c *gin.Context) {
	c.YAML(200, gin.H{"msg": "返回一个yaml"})
}

func returnProtoBuf(c *gin.Context) {
	// 略
}
