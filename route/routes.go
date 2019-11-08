package route

import (
	"alpaca_blog/controllers"

	"github.com/gin-gonic/gin"
)

func SetRoutes() (router *gin.Engine) {
	router = gin.Default() //获得路由实例
	router.Use(Middleware)
	router.GET("/ping", controllers.Pong) //短信验证码
	return
}

//Middleware Middleware
func Middleware(c *gin.Context) {
	//fmt.Println("this is a middleware!")
}
