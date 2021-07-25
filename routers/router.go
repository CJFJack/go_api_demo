package routers

import (
	"github.com/astaxie/beego"
	v1 "go_api_demo/controllers/api/v1"
)

func init() {

	// v1
	v1Test := beego.NewNamespace("v1", beego.NSAutoRouter(&v1.Test{}))
	beego.AddNamespace(v1Test)

}
