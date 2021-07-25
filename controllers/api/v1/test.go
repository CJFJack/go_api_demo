package v1

import (
	"github.com/astaxie/beego/logs"
	"go_api_demo/base/auth"
	"go_api_demo/base/response"
	"go_api_demo/forms"
)

type Test struct {
	auth.APIController
}

func (c *Test) Post() {
	if c.Ctx.Request.Method == "POST" {
		c.Data["json"] = response.Ok

		form := forms.TestPostForm{}

		if rawData, err := c.ParsePostForm(form); err == nil {
			logs.Info(rawData)
			// do something ...
		} else {
			logs.Error(err)
			c.Data["json"] = response.BadRequest
		}
	} else {
		c.Data["json"] = response.BadMethod
	}
}

func (c *Test) Get() {
	if c.Ctx.Request.Method == "GET" {
		c.Data["json"] = response.Ok

		//uuid := c.GetString("uuid")
		//logs.Info(uuid)

	} else {
		c.Data["json"] = response.BadMethod
	}
}
