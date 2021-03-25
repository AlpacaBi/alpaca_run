package route

import (
	"alpaca_blog/controllers"

	"github.com/gin-gonic/gin"
)

//InitRoutes 路由初始化
func InitRoutes() (router *gin.Engine) {
	router = gin.Default() //获得路由实例
	router.Use(Middleware) //中间件

	AI := router.Group("/ai")
	{
		AI.POST("/text", controllers.AIText)
		AI.POST("/image", controllers.AIImage)
	}

	router.GET("/github", controllers.Github) //测试！！！

	router.GET("/ping", controllers.Pong) //测试！！！
	return
}

//Middleware Middleware
func Middleware(c *gin.Context) {
	//fmt.Println("this is a middleware!")
}
