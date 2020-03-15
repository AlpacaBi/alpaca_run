package controllers

import (
	"alpaca_blog/config"
	"alpaca_blog/lib/ai"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"time"

	sjson "github.com/bitly/go-simplejson"
	"github.com/gin-gonic/gin"
)

// AIText 用于文字AI处理
func AIText(c *gin.Context) {
	text := c.Request.FormValue("text")
	robot := ai.NewXiaosiRobot(config.Current.Xiaosi.AppID)

	dirtyWord := [...]string{
		"你是不是性無能所以在網上自我高潮找存在感",
		"你爸今晚连夜庭院种枇杷树",
		"你爸妈真会生 好的东西都自己留着",
		"我真的很想帮你治疗，但我是一名家庭医生，而你是一个孤儿。",
		"咋滴你管那么多呢？收粪车从你家门前路过你都要拿勺子尝尝咸淡？",
		"以前只知道厕所能被屎堵了，今天竟然看见有人脑子被屎堵了",
		"你能讲文明，讲素质的时候，我想是我再次投胎做人的时候。",
		"上帝把智慧洒满人间，就你打了个伞。",
		"你的户口本就是一动物百科",
		"你是我见过的容量最大的铅笔盒了，装那么多笔你不累吗？",
		"你这种脑残程度我称之为珍稀物种",
		"和你接触的时间越长，我就越喜欢狗，狗永远是狗，人有时候不是人。",
		"现在网上冲浪不需要智商的吗",
		"我是灵长类动物，和家禽没什么好说的",
		"我留你狗命是因为我想保护动物毕竟你做只狗不容易",
		"请你不要用你的排泄器官对我说话，这是很不礼貌的，谢谢！",
		"阁下长得真是天生励志！",
		"如果长的丑也算残疾的话，那你就不用工作了。",
		"我把你挂到迎客松上喜迎八方贵客",
		"你的嘴简直能给农田施肥",
		"把我气死，你当孤儿？",
		"只希望你吃饭有人喂 走路有人推 夜夜缠绵于病却长命百岁 不能生育却儿孙满堂 寿比昙花福如母猪 ​​​​",
		"别跟我说话，我有洁癖。",
		"别对着我叫，我小时候被狗吓过。",
		"你走的那天，风很大，走得很痛苦，火化的 时候还诈尸，一直喊着没有死，最后用铁链绑着烧完的。火很旺，家属很坚强一个哭的都 没有，甚至有一个忍不住笑出声。运骨灰的路上车翻了 ，把骨灰盒摔了，刚要捧点灰，来了一辆洒水车，洒水车走后，众人只好把混着泥巴的骨灰，随便捧点埋了，不久之后该地方拆迁，上面开了一家迪厅，每天重复播放着今天是个好日子。",
		"这么喜欢装逼 下辈子当条内裤好了",
		"百度不一定可以搜得到你 但搜狗一定可以",
		"W D N M D",
		"刚好，我QQ农场缺条🐶，我看你不错今天下午就来上班吧",
		"以后不要再说你一无所有了，你不是还有病吗",
		"真羡慕这种身份证拿出来就是全家福的人",
		"你妈买菜必涨价",
		"你的出生是杜蕾斯的一封道歉信",
		"天工造物不测，怎么造出你这么个东西",
		"尘归尘，土归土，把你骨灰扬了都不配做PM2.5",
	}

	//ThWb41NmTDGuNeN3OdbeaPtp
	//89kM8rOGp4j7gd42UAsuzAF346LD4XBy
	client := &http.Client{}

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

	tt := info.Get("access_token").MustString("")

	url2 := "https://aip.baidubce.com/rest/2.0/solution/v1/text_censor/v2/user_defined?access_token=" + tt

	resp2, err := http.PostForm(url2, url.Values{"text": {text}})
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

	fmt.Println(info2)

	var msg string
	if info2.Get("conclusionType").MustInt() == 1 {
		msg = robot.GetXiaosiReplyMsg(text, "1")
	} else {

		msg = "<b>检测到你的发言带有恶意：</b><br/>"

		keyword := ""

		for i := 0; i < len(info2.Get("data").MustArray()); i++ {
			if info2.Get("data").GetIndex(i).Get("type").MustInt() == 12 {

				fmt.Println(info2.Get("data").GetIndex(i).Get("type").MustInt())
				subType := info2.Get("data").GetIndex(i).Get("subType").MustInt()

				probability, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", info2.Get("data").GetIndex(i).Get("hits").GetIndex(0).Get("probability").MustFloat64()*100), 64)

				for j := 0; j < len(info2.Get("data").GetIndex(i).Get("hits").GetIndex(0).Get("words").MustArray()); j++ {
					keyword = keyword + "【" + info2.Get("data").GetIndex(i).Get("hits").GetIndex(0).Get("words").GetIndex(j).MustString() + "】"
				}

				if subType == 5 {
					msg = msg + "<span style=\"color:red\">低俗辱骂："
				} else if subType == 2 {
					msg = msg + "<span style=\"color:yellow\">文本色情："
				} else if subType == 1 {
					msg = msg + "<span style=\"color:black\">暴恐违禁："
				} else if subType == 3 {
					msg = msg + "<span style=\"color:black\">政治敏感："
				}

				msg = msg + "(" + strconv.FormatFloat(probability, 'f', -1, 64) + "%)</span><br/>"

			}

		}

		if keyword != "" {
			msg = msg + "<b>关键词：</b><br/>" + keyword
		}

		rand.Seed(time.Now().UnixNano())
		msg = msg + "<br/><br/><br/>" + dirtyWord[rand.Intn(len(dirtyWord)+1)]

	}

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
