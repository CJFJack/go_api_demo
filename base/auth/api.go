package auth

import (
	"fmt"
	"github.com/astaxie/beego"
	"go_api_demo/base"
	"go_api_demo/base/response"
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
