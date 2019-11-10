package main

import (
	"alpaca_blog/config"
	"alpaca_blog/route"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	//全局设置环境，此为开发环境，线上环境为gin.ReleaseMode
	if config.Current.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	//读取配置文件
	config.LoadConfig()
	//配置api路由
	router := route.SetRoutes()
	//开始监听
	http.ListenAndServe(config.Current.Listen, router)
}
