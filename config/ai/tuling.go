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

// ReqPara 图灵机器人请求体JSON格式
type tulingReqData struct {
	ReqType    int32 `json:"reqType"`
	Perception struct {
		InputText struct {
			Text string `json:"text"`
		} `json:"inputText"`
	} `json:"perception"`
	UserInfo struct {
		APIKey string `json:"apiKey"`
		UserID string `json:"userId"`
	} `json:"userInfo"`
}

// post 给Robot对象封装好post
func (r *Robot) post(url string, para []byte) ([]byte, error) {
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

// Robot 机器人对象结构体
type Robot struct {
	APIkey string
	APIURL string
	rnet   *http.Client
}

// NewRobot 初始化机器人
func NewRobot(apikey string) *Robot {
	return &Robot{
		APIkey: apikey,
		APIURL: config.Current.TulingAPIURL,
		rnet:   &http.Client{},
	}
}

// GetReplyMsg 正式调用图灵机器人接口
func (r *Robot) GetReplyMsg(text string, userid string) string {
	var para tulingReqData
	para.ReqType = 0
	para.Perception.InputText.Text = text
	para.UserInfo.APIKey = r.APIkey
	para.UserInfo.UserID = userid
	jsondata, _ := json.Marshal(para)

	body, err := r.post(r.APIURL, jsondata)
	if err != nil {
		fmt.Println("get tuling reply err:", err)
		return ""
	}

	info, err := sjson.NewJson(body)
	if err != nil {
		fmt.Println("get tuling reply parse para err:", err)
		return ""
	}

	msg := info.Get("results").GetIndex(0).Get("values").Get("text").MustString("")
	return msg
}
