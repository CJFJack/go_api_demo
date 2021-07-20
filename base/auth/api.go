package auth

import (
	"auto_tuomin_agent/base"
	"auto_tuomin_agent/base/response"
	"fmt"
	"github.com/astaxie/beego"
)

type APIController struct {
	base.BaseController
}

func (c *APIController) Prepare() {
	c.EnableXSRF = false

	token := fmt.Sprintf("Token %s", beego.AppConfig.DefaultString("api::token", ""))
	headerToken := c.Ctx.Input.Header("Authorization")

	if token != headerToken {
		c.Data["json"] = response.UnAuthorization
		c.ServeJSON()
	}
}
