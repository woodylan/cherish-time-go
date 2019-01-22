package timeComtroller

import (
	"cherish-time-go/controllers"
	"github.com/astaxie/beego/validation"
	"cherish-time-go/logic/weapp/time"
)

type TimeDeleteController struct {
	controllers.Controller
}

func (c *TimeDeleteController) Delete() {
	inputData := struct {
		Id string `data:"id" valid:"Required"`
	}{}
	c.GetData(&inputData)
	c.Valid(&inputData)

	valid := validation.Validation{}
	valid.Valid(&inputData)

	timeLogic := timeLogic.TimeLogic{}
	timeLogic.Delete(c.Ctx, inputData.Id)

	c.Data["json"] = controllers.RetData{0, "success", make([]string, 0)}
	c.ServeJSON()
}
