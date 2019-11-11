package controllers

import (
	"github.com/gin-gonic/gin"
)

//Pong 测试用api，本项目第一个api
func Pong(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
	return
}
