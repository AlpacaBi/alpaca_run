package config

import (
	"flag"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"

	jsoniter "github.com/json-iterator/go"
	"hz.com/golib/utils"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type configuration struct {
	Listen  string `json:"listen"`
	Debug   bool   `json:"debug"`
	LogFile string `json:"logFile"` //日志
}

//Current the current configuration
var Current configuration

//Init init the configuration
func init() {
	Current = configuration{
		":9527",
		true,
		"alpaca_blog.log",
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
		logPath = path.Join(utils.GetRootPath(), logPath)
	}

	if !Current.Debug { //如果是调试模式，则直接输出到标准控制输出
		file, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			panic(err)
		}
		log.SetOutput(file)
	}
}
