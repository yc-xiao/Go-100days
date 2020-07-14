package routers

import (
	"BeegoDemo2/controllers"
	"github.com/astaxie/beego"
)

func init() {

	//获取用户信息
	//get
	beego.Router("/getUserInfo", &controllers.CustomController{}, "GET:GetUserInfo")
	//post
	beego.Router("/registerInfo", &controllers.CustomController{}, "POST:RegisterInfo")
}
