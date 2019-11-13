package controllers

import (
	"alpaca_blog/config"
	"alpaca_blog/lib/ai"

	"github.com/gin-gonic/gin"
)

//AIText 用于文字AI处理
func AIText(c *gin.Context) {
	text := c.Request.FormValue("text")
	robot := ai.NewRobot(config.Current.TulingAPIKey)
	msg := robot.GetReplyMsg(text, "1")
	c.JSON(200, gin.H{
		"status":  "ok",
		"message": msg,
	})
	return
}

//AIImage 用于图像AI处理
func AIImage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
	return
}
