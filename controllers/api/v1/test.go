package v1

import (
	"auto_tuomin_agent/base/auth"
	"auto_tuomin_agent/base/response"
	"auto_tuomin_agent/forms"
	"github.com/astaxie/beego/logs"
)

type Test struct {
	auth.APIController
}

func (c *Test) Post() {
	if c.Ctx.Request.Method == "POST" {
		c.Data["json"] = response.Ok

		form := forms.TestPostForm{}

		if rawData ,err := c.ParsePostForm(form); err == nil {
			logs.Info(rawData)
			// do something ...
		} else {
			logs.Error(err)
			c.Data["json"] = response.BadRequest
		}
	}
	c.Data["json"] = response.BadMethod
}

func (c *Test) Get() {
	if c.Ctx.Request.Method == "GET" {
		c.Data["json"] = response.Ok

		//uuid := c.GetString("uuid")
		//logs.Info(uuid)

	}
	c.Data["json"] = response.BadMethod
}
