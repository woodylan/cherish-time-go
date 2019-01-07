package timeComtroller

import (
	"cherish-time-go/controllers"
	"github.com/astaxie/beego/validation"
)

type TimeDetailController struct {
	controllers.Controller
	InputData InputData
}

type InputData struct {
	Id string `data:"id" valid:"IP"`
}

func (c *TimeDetailController) Detail() {
	inputData := InputData{}
	c.GetData(&inputData)
	c.Valid(&inputData)

	valid := validation.Validation{}
	valid.Valid(&inputData)

	c.Data["json"] = inputData.Id
	c.ServeJSON()
}
