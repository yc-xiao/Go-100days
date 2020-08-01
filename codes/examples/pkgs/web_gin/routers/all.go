package routers

import "github.com/gin-gonic/gin"

func Register(r *gin.Engine) {
	simple := r.Group("/simple")
	{
		simple.GET("/:name/*action", urlParam)
		simple.GET("/", getQuery)
		simple.POST("/", postForm)
		simple.POST("/uploadFile/", uploadFile)
	}

	data := r.Group("/data")
	{
		data.POST("loginByJson", loginJsonByjson)
		data.POST("loginByForm", loginJsonByform)
	}

	reponse := r.Group("/reponse")
	{
		reponse.GET("text", returnText)
		reponse.GET("json", returnJson)
		reponse.GET("xml", returnXml)
		reponse.GET("yaml", returnYaml)

	}
	middle := r.Group("/middle")
	{
		middle.GET("/login", login)
		middle.GET("/logout", logout)
		middle.GET("/home", CheckCookies(), home)
	}
}
