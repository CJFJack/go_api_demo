package main

import (
	_ "auto_tuomin_agent/routers"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run(fmt.Sprintf(":%s", beego.AppConfig.DefaultString("Port", "9801")))
}
