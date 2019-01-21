package userController

import (
	"cherish-time-go/controllers"
	"github.com/astaxie/beego/validation"
	"cherish-time-go/logic/weapp/account"
)

type CheckAuthController struct {
	controllers.Controller
}

func (c *CheckAuthController) CheckAuth() {
	inputData := struct {
		Auth string `data:"auth" valid:"Required"`
	}{}
	c.GetData(&inputData)
	c.Valid(&inputData)

	valid := validation.Validation{}
	valid.Valid(&inputData)

	logic := accountLogic.AccountLogic{}
	ret := logic.CheckAuth(c.Ctx, inputData.Auth)

	c.Data["json"] = controllers.RetData{0, "success", ret}
	c.ServeJSON()
}
