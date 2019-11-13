# Alpaca Blog

## 简介
本项目是一个个人博客，开发的技术栈和工具有
- <b style="color:#2b7489">TypeScript</b>
- <b style="color:rgb(79, 192, 141)">Vue.js</b>
- <b style="color:rgb(0, 173, 216)">Golang</b>
- Gin
- Redis
- MySQL
- Graphql（暂时没上）

## 如何运行
注意：本项目目前还在开发中，大部分功能尚未开发完成

### 事先做好准备：
- 安装Node.js
- 安装Go和配置相关环境变量
- 装好redis和mysql
- 安装以下Go库（go get一把梭就好）
  - github.com/gin-gonic/gin
  - github.com/go-redis/redis
  - github.com/go-sql-driver/mysql
  - github.com/jmoiron/sqlx
  - github.com/json-iterator/go
  - github.com/bitly/go-simplejson

### 前端：
本项目的前端部分在根目录的`views`文件夹下，所以要先进入`views`文件夹  
本前端项目是由vue-cli3搭建的，所以也是常规方式：  
安装npm包 `npm install`  
开发模式  `npm run serve`  
打包     `npm run build`
当你打包好后，把打包好的文件丢到linux服务器，再用nginx配置下就好了，怎么配置自己google

### 后端：
1. 把项目复制到Go目录的src目录下，直接`go build main.go`就能打包出可执行文件`main`了  
2. 因为我用的linux电脑，所以可以直接这么来，如果你用的是window电脑，要交叉编译改参数，这个你google就好，很简单的  
3. 把可执行文件`main`和`config.json`，直接塞进你的linux服务器里，直接`nohup ./main &`就能跑起来了（他会自动读config.json配置的  

如果想开发中调试，我用的vscode，可以给你参考下launch.json，args根据你电脑文件目录来
```json
{
    "version": "1.0.0",
    "configurations": [        
        {
            "name": "Launch",
            "type": "go",
            "request": "launch",
            "mode": "debug",
            "program": "${workspaceRoot}",
            "env": {},
            "args": ["-config", "/home/alpaca/GoEnv/src/alpaca_blog/config.json"]
        }
    ]
}
```
### config.json
可以更改config.json配置来动态更改go项目参数配置，重启你的go web服务就能生效
```json
{
    "listen":":9527",
    "debug": true,
    "logFile": "alpaca_bi.log",
    "mysql": {
        "host": "localhost",
        "port": 3306,
        "user": "root",
        "password": "填你自己的mysql密码",
        "db": "alpacablog",
        "dbprefix": "alpaca_blog_"
    },
      "redis": {
        "host": "localhost",
        "port": 6379,
        "auth": "填你自己的redis密码,一般没特地设置的话密码为空",
        "db": 0,
        "poolSize": 300
    },
    "tulingAPIKey": "填你自己的图灵机器人APIKey，去官网注册就能免费拿",
    "tulingAPIURL":"http://openapi.tuling123.com/openapi/api/v2"
}
```

## 联系我
如果你遇到问题，或者想交流技术，可以加我微信或者电报
<p align="left">
	<img src="https://alpaca.cdn.bcebos.com/wechat.jpg" alt="wechat"  width="250">
</p>

