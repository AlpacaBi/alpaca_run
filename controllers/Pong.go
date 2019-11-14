package controllers

import (
	"alpaca_blog/lib/mysql"
	"alpaca_blog/lib/redis"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

//Pong 测试用api，本项目第一个api
func Pong(c *gin.Context) {
	db := mysql.GetDB()

	// 测试mysql
	updateSQL := fmt.Sprintf("update #__ping set count=count+1 where id = %d", 1)
	_, err := db.Exec(mysql.Prefix(updateSQL))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	seleceSQL := fmt.Sprintf("select count from #__ping where id = %d", 1)
	var count int
	err2 := db.Get(&count, mysql.Prefix(seleceSQL))
	if err2 != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "error",
			"error":  err2.Error(),
		})
		return
	}

	// 测试redis
	redis.Set("ping", time.Now().Format("2006-01-02 15:04:05"), 2*time.Hour)
	value, err := redis.Get("ping")
	if err != nil {
		fmt.Println(err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": value,
		"count":   count,
	})
	return
}
