package timeLogic

import (
	"cherish-time-go/models/Time"
	"fmt"
)

type TimeLogic struct {
}

type TimeDetail struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Type       uint8   `json:"type"`
	Color      string `json:"color"`
	Data       string `json:"data"`
	Days       string `json:"days"`
	Remark     string `json:"remark"`
	CreateTime string `json:"createTime"`
}

func (this *TimeLogic) GetDetail(id string) (timeDetail TimeDetail) {
	model := TimeModel.GetById(id)

	fmt.Println(model)

	timeDetail = this.renderDetail(model)

	return
}

func (this *TimeLogic) renderDetail(model TimeModel.Time) (timeDetail TimeDetail) {
	timeDetail.Id = model.Id
	timeDetail.Name = model.Name
	timeDetail.Type = model.Type
	return
}
