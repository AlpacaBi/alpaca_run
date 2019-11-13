package controllers

import (
	"alpaca_blog/config/redis"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

//Pong 测试用api，本项目第一个api
func Pong(c *gin.Context) {
	redis.Set("ping", time.Now().Format("2006-01-02 15:04:05"), 2*time.Hour)
	value, err := redis.Get("ping")
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(200, gin.H{
		"message": value,
	})
	return
}
