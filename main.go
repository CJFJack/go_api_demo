package main

import (
	_ "auto_tuomin_agent/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run(":9802")
}
