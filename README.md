# Alpaca Blog

## 简介
本项目是一个个人博客，开发的技术栈有
- typescript
- vue2.6
- golang
- gin
- graphql
- mysql

## 如何运行
注意：本项目目前还在开发中，大部分功能尚未开发完成
### 前端：
在此之前，你需要在你电脑装node.js,至于怎么装自己google吧  
本前端项目在根目录的`views`文件夹下，所以要先进入`views`文件夹  
本前端项目是由vue-cli3搭建的，所以也是常规方式：  
安装npm包 `npm install`  
开发模式  `npm run serve`  
打包     `npm run build`
### 后端：
本项目用的Golang作为web后台语言，所以你需要在你的电脑安装Go和配置相关环境变量等等，怎么配置自己google吧  
把项目复制到Go目录的src目录下，直接`go build main.go`就能打包出可执行文件了。  
如果想开发中调试，我用的vscode，可以给你参考下launch.json
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
