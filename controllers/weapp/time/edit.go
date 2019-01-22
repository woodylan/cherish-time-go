package timeComtroller

import (
	"cherish-time-go/controllers"
	"github.com/astaxie/beego/validation"
	"cherish-time-go/logic/weapp/time"
)

type TimeEditController struct {
	controllers.Controller
}

func (c *TimeEditController) Edit() {
	inputData := struct {
		Id     string   `data:"name" valid:"Required"`
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
	logicRet := timeLogic.Edit(c.Ctx, inputData.Id, inputData.Name, inputData.Color, inputData.Date, inputData.Remark)

	c.Data["json"] = controllers.RetData{0, "success", logicRet}
	c.ServeJSON()
}
