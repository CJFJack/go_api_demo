package main

import (
	"fmt"
	"github.com/astaxie/beego"
	_ "go_api_demo/routers"
)

func main() {
	beego.Run(fmt.Sprintf(":%s", beego.AppConfig.DefaultString("Port", "9801")))
}
