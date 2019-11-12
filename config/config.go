package config

import (
	"alpaca_blog/config/redis"
	"flag"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"

	"alpaca_blog/config/mysql"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type configuration struct {
	Listen       string       `json:"listen"`
	Debug        bool         `json:"debug"`
	LogFile      string       `json:"logFile"` //日志
	Mysql        mysql.Config `json:"mysql"`
	Redis        redis.Config `json:"redis"`
	TulingAPIKey string       `json:"tulingAPIKey"` //图灵机器人APIKey
	TulingAPIURL string       `json:"tulingAPIURL"` //图灵机器人接口URL
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
			"", //"127.0.0.1"
			0,  //3306
			"", //"root"
			"", //""
			"", //"alpacablog"
			"", //"alpaca_blog_"
			0,  //40
			0,  //2
			0,  //30
		},
		redis.Config{
			"", //"127.0.0.1"
			0,  //6379
			"", //""
			0,  //0
			0,  //300
		},
		"",
		"",
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
