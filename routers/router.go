package routers

import (
	v1 "auto_tuomin_agent/controllers/api/v1"
	"github.com/astaxie/beego"
)

func init() {

	// v1
	v1Test := beego.NewNamespace("v1", beego.NSAutoRouter(&v1.Test{}))
	beego.AddNamespace(v1Test)

}
