package timeComtroller

import (
	"cherish-time-go/controllers"
	"github.com/astaxie/beego/validation"
	"cherish-time-go/logic/weapp/time"
)

type TimeListController struct {
	controllers.Controller
}

func (c *TimeListController) List() {
	inputData := struct {
		PerPage     int `data:"perPage" valid:"Min(1)"`
		CurrentPage int `data:"currentPage" valid:"Min(1)"`
	}{}
	c.GetData(&inputData)
	c.Valid(&inputData)

	valid := validation.Validation{}
	valid.Valid(&inputData)

	timeLogic := timeLogic.TimeLogic{}
	logicRet := timeLogic.GetList(c.Ctx, inputData.PerPage, inputData.CurrentPage)

	c.Data["json"] = controllers.RetData{0, "success", logicRet}
	c.ServeJSON()
}
