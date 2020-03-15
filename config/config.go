package config

import (
	"alpaca_blog/lib/mysql"
	"alpaca_blog/lib/redis"
	"flag"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type configuration struct {
	Listen  string       `json:"listen"`
	Debug   bool         `json:"debug"`
	LogFile string       `json:"logFile"` //日志
	Mysql   mysql.Config `json:"mysql"`
	Redis   redis.Config `json:"redis"`
	Tuling  Tuling       `json:"tuling"`  //图灵机器人
	Xiaosi  Xiaosi       `json:"xiaosi"`  //小思机器人
	BaiduAI BaiduAI      `json:"baiduAI"` //小思机器人

}

//Tuling 图灵机器人参数
type Tuling struct {
	APIKey string `json:"apiKey"` //图灵机器人APIKey
	APIURL string `json:"apiURL"` //图灵机器人接口URL
}

//Xiaosi 小思机器人参数
type Xiaosi struct {
	AppID  string `json:"appID"`  //小思机器人APIKey
	APIURL string `json:"apiURL"` //小思机器人接口URL
}

//BaiduAI 百度AI机器人参数
type BaiduAI struct {
	APIKey    string `json:"apiKey"`    //百度AI APIKey
	SecretKey string `json:"secretKey"` //百度AI SecretKey
}

//Current the current configuration
var Current configuration

//Init init the configuration
func init() {
	Current = configuration{
		":9527",
		true,
		"alpaca_blog.log",
		mysql.Config{
			"",
			0,
			"",
			"",
			"",
			"",
			0,
			0,
			0,
		},
		redis.Config{
			"",
			0,
			"",
			0,
			0,
		},
		Tuling{
			"",
			"",
		},
		Xiaosi{
			"",
			"",
		},
		BaiduAI{
			"",
			"",
		},
	}

	basePath, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(basePath)
	configFile := filepath.Join(filepath.Dir(path), "config.json")
	if len(os.Args) > 1 {
		configFile = os.Args[1]
	}
	file, err := os.Open(configFile)

	if err != nil {
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Current)
	if err != nil {
		log.Print(err)
		return
	}
}

//LoadFromFile load config from file
func LoadFromFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	return decoder.Decode(&Current)
}

func GetRootPath() string {
	basePath, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(basePath)
	return filepath.Dir(path)
}

func LoadConfig() {
	configFile := flag.String("config", "", "The config file path.")
	flag.Parse()

	// load config
	if len(*configFile) > 0 {
		log.Printf("load config from file: %s\n", *configFile)
		LoadFromFile(*configFile)
	}

	//setup log store path
	logPath := Current.LogFile
	if !filepath.IsAbs(logPath) {
		logPath = path.Join(GetRootPath(), logPath)
	}

	if !Current.Debug { //如果是调试模式，则直接输出到标准控制输出
		file, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			panic(err)
		}
		log.SetOutput(file)
	}
}
