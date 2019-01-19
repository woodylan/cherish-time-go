package userController

import (
	"cherish-time-go/controllers"
	"github.com/astaxie/beego/validation"
	"cherish-time-go/logic/weapp/account"
)

type UserLoginController struct {
	controllers.Controller
}

func (c *UserLoginController) Login() {
	inputData := struct {
		Code          string `data:"code" valid:"Required"`
		Iv            string `data:"iv" valid:"Required"`
		EncryptedData string `data:"encryptedData" valid:"Required"`
	}{}
	c.GetData(&inputData)
	c.Valid(&inputData)

	valid := validation.Validation{}
	valid.Valid(&inputData)

	logic := accountLogic.AccountLogic{}
	ret := logic.Login(c.Ctx, inputData.Code, inputData.Iv, inputData.EncryptedData)

	c.Data["json"] = controllers.RetData{0, "success", ret}
	c.ServeJSON()
}
