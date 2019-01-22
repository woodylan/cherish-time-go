package timeLogic

import (
	"cherish-time-go/models/Time"
	"cherish-time-go/modules/util"
	"time"
	"cherish-time-go/define/common"
	"cherish-time-go/controllers"
	"cherish-time-go/define/retcode"
	"github.com/astaxie/beego/context"
	"cherish-time-go/global"
	"cherish-time-go/models/Sentence"
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

type TimeListDetail struct {
	Id         string                 `json:"id"`
	Name       string                 `json:"name"`
	Type       uint8                  `json:"type"`
	Color      []string               `json:"color"`
	Date       string                 `json:"date"`
	Days       int64                  `json:"days"`
	Remark     string                 `json:"remark"`
	Sentence   SentenceModel.Sentence `json:"sentence"`
	CreateTime int64                  `json:"createTime"`
}

func (this *TimeLogic) GetDetail(c *context.Context, id string) (timeDetail TimeDetail) {
	LoginUserInfo := global.LoginUserInfo

	model, err := TimeModel.GetById(id)
	if err != nil || model.UserId != LoginUserInfo.UserId {
		util.ThrowApi(c, retcode.ERR_OBJECT_NOT_FOUND, "找不到数据"+err.Error())
		return
	}

	timeDetail = this.renderDetail(model)

	return
}

func (this *TimeLogic) GetList(c *context.Context, perPage, currentPage int) (page controllers.Page) {
	LoginUserInfo := global.LoginUserInfo

	models, sumCount, err := TimeModel.GetByPage(LoginUserInfo.UserId, perPage, currentPage)
	if err != nil {
		util.ThrowApi(c, retcode.ERR_OBJECT_NOT_FOUND, "找不到数据"+err.Error())
		return
	}

	//SentenceModel.GetRand(perPage)
	sentences, err := SentenceModel.GetRand(perPage)
	//fmt.Println(sentences)

	page.RendPage(sumCount, perPage, currentPage)
	this.renderList(&page, models, sentences, sumCount)

	return
}

func (this *TimeLogic) Create(c *context.Context, name, color, date, remark string) {
	timeType := getTypeByDate(date)
	LoginUserInfo := global.LoginUserInfo

	_, ok := TimeModel.AddNew(name, LoginUserInfo.UserId, timeType, date, color, remark)
	if !ok {
		util.ThrowApi(c, retcode.ERR_WRONG_MYSQL_OPERATE, "新建记录失败")
	}

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

func (this *TimeLogic) renderList(page *controllers.Page, models []TimeModel.Time, sentences []SentenceModel.Sentence, sumCount int) (*controllers.Page) {
	nowTimeUnix := time.Now().Unix()

	list := make([]TimeListDetail, 0)
	for key, model := range models {
		timeDetail := TimeListDetail{}
		dateTime, _ := time.Parse("20060102", model.Date)
		dateTimeUnix := dateTime.Unix()

		days := int64(0)
		if model.Type == common.TIME_TYPE_DESC {
			days = util.DaysDiff(nowTimeUnix, dateTimeUnix)
		} else if model.Type == common.TIME_TYPE_ASC {
			days = util.DaysDiff(dateTimeUnix, nowTimeUnix)
		}

		color := []string{};
		if len(model.Color) > 0 {
			util.JsonDecode(model.Color, &color)
		}

		timeDetail.Id = model.Id
		timeDetail.Name = model.Name
		timeDetail.Type = model.Type
		timeDetail.Color = color
		timeDetail.Date = model.Date
		timeDetail.Days = days
		timeDetail.Remark = model.Remark
		timeDetail.Sentence = sentences[key]
		timeDetail.CreateTime = model.CreatedAt.Unix()
		
		list = append(list, timeDetail)
	}

	page.List = list
	return page
}

//通过日期与当前日期比较，得到时间类型（正计时还是倒计时）
func getTypeByDate(targetDate string) uint8 {
	nowDate := time.Now().Format("20060102")
	if targetDate > nowDate {
		//目标日
		return common.TIME_TYPE_DESC
	}

	return common.TIME_TYPE_ASC
}
