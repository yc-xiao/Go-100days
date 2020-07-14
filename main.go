package main
import "github.com/astaxie/beego"

func main() {
	beego.Info("第一个beego案例")
	beego.Run("localhost:8080")
}
