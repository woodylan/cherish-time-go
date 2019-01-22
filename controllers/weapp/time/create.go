package timeComtroller

import (
	"cherish-time-go/controllers"
	"github.com/astaxie/beego/validation"
	"cherish-time-go/logic/weapp/time"
)

type TimeCreateController struct {
	controllers.Controller
}

func (c *TimeCreateController) Create() {
	inputData := struct {
		Name   string   `data:"name" valid:"Required"`
		Color  []string `data:"color" valid:"Required"`
		Date   string   `data:"date" valid:"Required"`
		Remark string   `data:"remark" valid:""`
	}{}

	c.GetData(&inputData)
	c.Valid(&inputData)

	valid := validation.Validation{}
	valid.Valid(&inputData)

	timeLogic := timeLogic.TimeLogic{}
	timeLogic.Create(c.Ctx, inputData.Name, inputData.Color, inputData.Date, inputData.Remark)

	c.Data["json"] = controllers.RetData{0, "success", ""}
	c.ServeJSON()
}
