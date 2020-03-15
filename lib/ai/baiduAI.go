package ai

import (
	"alpaca_blog/config"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	sjson "github.com/bitly/go-simplejson"
)

// baiduAPI 百度AI接口
type baiduAPI struct {
	AccessToken   string //获取accesstoken地址
	TextCensor    string //文本审核
	ImageClassify string //图像识别
}

// baiduAI 百度AI对象结构体
type baiduAI struct {
	APIkey    string
	SecretKey string
	BaiduAPI  *baiduAPI
	rnet      *http.Client
}

// NewBaiduAI 初始化百度AI
func NewBaiduAI() *baiduAI {
	return &baiduAI{
		APIkey:    config.Current.BaiduAI.APIKey,
		SecretKey: config.Current.BaiduAI.SecretKey,
		BaiduAPI: &baiduAPI{
			AccessToken:   "https://aip.baidubce.com/oauth/2.0/token",
			TextCensor:    "https://aip.baidubce.com/rest/2.0/solution/v1/text_censor/v2/user_defined",
			ImageClassify: "https://aip.baidubce.com/rest/2.0/image-classify/v2/advanced_general",
		},
		rnet: &http.Client{},
	}
}

// getAccessToken 获取AccessToken
func (baiduAI *baiduAI) getAccessToken() (accessToken string, err error) {
	accessTokenURL := fmt.Sprintf("%s?grant_type=client_credentials&client_id=%s&client_secret=%s", baiduAI.BaiduAPI.AccessToken, baiduAI.APIkey, baiduAI.SecretKey)

	req, err := http.NewRequest("GET", accessTokenURL, nil)
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "appliction/json;charset=utf-8")
	resp, err := baiduAI.rnet.Do(req)
	if err != nil {
		return
	}

	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	info, err := sjson.NewJson(body)
	if err != nil {
		return
	}

	accessToken = info.Get("access_token").MustString("")

	return
}

// ImageClassify 图像识别
func (baiduAI *baiduAI) ImageClassify(base64URL string) (msg *sjson.Json, err error) {

	accessToken, err := baiduAI.getAccessToken()
	if err != nil {
		fmt.Println(err)
		return
	}

	url2 := fmt.Sprintf("%s?access_token=%s", baiduAI.BaiduAPI.ImageClassify, accessToken)

	resp2, err := http.PostForm(url2, url.Values{"image": {base64URL}, "baike_num": {"1"}})
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp2.Body.Close()
	body, _ := ioutil.ReadAll(resp2.Body)

	msg, err = sjson.NewJson(body)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

// ImageClassify 文本审核
func (baiduAI *baiduAI) TextCensor(text string) (msg *sjson.Json, err error) {

	accessToken, err := baiduAI.getAccessToken()
	if err != nil {
		fmt.Println(err)
		return
	}

	url2 := fmt.Sprintf("%s?access_token=%s", baiduAI.BaiduAPI.TextCensor, accessToken)

	resp2, err := http.PostForm(url2, url.Values{"text": {text}})
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp2.Body.Close()
	boo2, _ := ioutil.ReadAll(resp2.Body)

	msg, err = sjson.NewJson(boo2)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}
