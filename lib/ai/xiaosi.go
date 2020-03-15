package ai

import (
	"alpaca_blog/config"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	sjson "github.com/bitly/go-simplejson"
)

// ReqPara 小思机器人请求体JSON格式
type xiaosiReqData struct {
	Spoken string `json:"spoken"`
	AppID  string `json:"appid"`
	UserID string `json:"userid"`
}

// post 给Robot对象封装好post
func (r *xiaosiRobot) post(url string, para []byte) ([]byte, error) {
	body := bytes.NewBuffer(para)
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "appliction/json;charset=utf-8")
	resp, err := r.rnet.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

// Robot 小思机器人对象结构体
type xiaosiRobot struct {
	AppID  string
	APIURL string
	rnet   *http.Client
}

// NewXiaosiRobot 初始化小思机器人
func NewXiaosiRobot(appid string) *xiaosiRobot {
	return &xiaosiRobot{
		AppID:  appid,
		APIURL: config.Current.Xiaosi.APIURL,
		rnet:   &http.Client{},
	}
}

// GetReplyMsg 正式调用小思机器人接口
func (r *xiaosiRobot) GetXiaosiReplyMsg(text string, userid string) string {
	var para xiaosiReqData
	para.Spoken = text
	para.AppID = r.AppID
	para.UserID = userid
	jsondata, _ := json.Marshal(para)

	body, err := r.post(r.APIURL, jsondata)
	if err != nil {
		fmt.Println("get xiaosi reply err:", err)
		return ""
	}

	info, err := sjson.NewJson(body)
	if err != nil {
		fmt.Println("get xiaosi reply parse para err:", err)
		return ""
	}

	msg := info.Get("data").Get("info").Get("text").MustString("")
	return msg
}
