package controllers

import (
	"alpaca_blog/config"
	"alpaca_blog/lib/ai"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	sjson "github.com/bitly/go-simplejson"
	"github.com/gin-gonic/gin"
)

// AIText 用于文字AI处理
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

// AIImage 用于图像识别AI处理
func AIImage(c *gin.Context) {
	client := &http.Client{}
	base64URL := c.Request.FormValue("imgBase64")

	url1 := "https://aip.baidubce.com/oauth/2.0/token?grant_type=client_credentials&client_id=xv8D6Pe7OTZ83y60p9GGIE73&client_secret=kWGQRsxOCDZZPcTXcS7VTCgrnClCAsln"

	req, err := http.NewRequest("GET", url1, nil)
	if err != nil {
		fmt.Println(1)
		fmt.Println(err)
	}

	req.Header.Set("Content-Type", "appliction/json;charset=utf-8")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(2)
		fmt.Println(err)
	}

	boo, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	info, err := sjson.NewJson(boo)
	if err != nil {
		fmt.Println(3)
		fmt.Println(err)
	}

	msg := info.Get("access_token").MustString("")

	url2 := "https://aip.baidubce.com/rest/2.0/image-classify/v2/advanced_general?access_token=" + msg

	resp2, err := http.PostForm(url2, url.Values{"image": {base64URL}, "baike_num": {"1"}})
	if err != nil {
		fmt.Println(2)
		fmt.Println(err)
	}
	defer resp2.Body.Close()
	boo2, _ := ioutil.ReadAll(resp2.Body)

	info2, err := sjson.NewJson(boo2)
	if err != nil {
		fmt.Println(3)
		fmt.Println(err)
	}

	c.JSON(200, gin.H{
		"status":  "ok",
		"message": info2,
	})
	return
}
