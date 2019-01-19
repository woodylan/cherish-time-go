package timeLogic

import (
	"cherish-time-go/models/Time"
	"cherish-time-go/modules/util"
	"time"
	"cherish-time-go/define/common"
	"github.com/astaxie/beego"
	"cherish-time-go/controllers"
)

type TimeLogic struct {
}

type TimeDetail struct {
	Id         string   `json:"id"`
	Name       string   `json:"name"`
	Type       uint8    `json:"type"`
	Color      []string `json:"color"`
	Date       string   `json:"date"`
	Days       int64    `json:"days"`
	Remark     string   `json:"remark"`
	CreateTime int64    `json:"createTime"`
}

func (this *TimeLogic) GetDetail(id string) (timeDetail TimeDetail) {
	model, err := TimeModel.GetById(id)
	if err != nil {
		beego.BeeLogger.Error("Error finding user with id %s: %v", id, err.Error())
	}

	timeDetail = this.renderDetail(model)

	return
}

func (this *TimeLogic) GetList(perPage, currentPage int) (page controllers.Page) {
	models, sumCount, err := TimeModel.GetByPage("922611e8adad83fc", perPage, currentPage)
	if err != nil {
		beego.BeeLogger.Error("Error get users error: %v", err.Error())
	}

	page.RendPage(sumCount, perPage, currentPage)
	this.renderList(&page, models, sumCount)

	return
}

func (this *TimeLogic) renderDetail(model TimeModel.Time) (timeDetail TimeDetail) {
	color := []string{};
	if len(model.Color) > 0 {
		util.JsonDecode(model.Color, &color)
	}

	nowTimeUnix := time.Now().Unix()
	dateTime, _ := time.Parse("20060102", model.Date)
	dateTimeUnix := dateTime.Unix()
	days := int64(0)
	if model.Type == common.TIME_TYPE_DESC {
		days = util.DaysDiff(nowTimeUnix, dateTimeUnix)
	} else if model.Type == common.TIME_TYPE_ASC {
		days = util.DaysDiff(dateTimeUnix, nowTimeUnix)
	}

	timeDetail.Id = model.Id
	timeDetail.Name = model.Name
	timeDetail.Type = model.Type
	timeDetail.Color = color
	timeDetail.Date = model.Date
	timeDetail.Days = days
	timeDetail.Remark = model.Remark
	timeDetail.CreateTime = model.CreatedAt.Unix()
	return
}

func (this *TimeLogic) renderList(page *controllers.Page, models []TimeModel.Time, sumCount int) (*controllers.Page) {
	list := make([]TimeDetail, 0)
	for _, val := range models {
		modelValue := this.renderDetail(val)
		//todo 美句
		list = append(list, modelValue)
	}

	page.List = list
	return page
}
