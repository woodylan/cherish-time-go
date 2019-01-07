package timeComtroller

import (
	"cherish-time-go/controllers"
)

type TimeDetailController struct {
	controllers.Controller
}

type InputData struct {
	Id   string `data:"id"`
}

func (c *TimeDetailController) Detail() {
	inputData := InputData{}
	c.GetData(&inputData)

	c.Data["json"] = inputData.Id
	c.ServeJSON()
}
