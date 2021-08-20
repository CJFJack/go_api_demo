package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/beego/beego/v2/core/logs"
	_ "go_api_demo/routers"
)

func initLog() {
	logs.SetLogger(logs.AdapterFile, fmt.Sprintf(`{"filename":"%s","level":%d,"maxlines":%d,"maxsize":%d,"daily":%t,"maxdays":%d,"color":%t}`,
		beego.AppConfig.DefaultString("log::Filename", "logs/demo.log"),
		beego.AppConfig.DefaultInt("log::Level", 7),
		beego.AppConfig.DefaultInt("log::MaxLines", 0),
		beego.AppConfig.DefaultInt("log::MaxSize", 0),
		beego.AppConfig.DefaultBool("log::Daily", true),
		beego.AppConfig.DefaultInt("log::MaxDays", 7),
		beego.AppConfig.DefaultBool("log::Color", true),
	))
	if beego.AppConfig.DefaultString("Env", "dev") != "PRO" {
		logs.SetLogger("console")
	}
	logs.EnableFuncCallDepth(true)
	logs.SetLevel(logs.LevelDebug)
	logs.Async()
}

func main() {
	// 初始化日志
	initLog()
	// 运行
	beego.Run(fmt.Sprintf(":%s", beego.AppConfig.DefaultString("Port", "9801")))
}
