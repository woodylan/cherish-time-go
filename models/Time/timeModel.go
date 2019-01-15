package TimeModel

import (
	"github.com/astaxie/beego/orm"
)

// Model Struct
type Time struct {
	Id         string `orm:"column(id);pk"`
	Name       string
	UserId     string
	Type       uint8
	Date       string
	Color      string
	Remark     string
	CreateTime int64
}

func (a *Time) TableName() string {
	return "time"
}

func GetById(id string) (Time) {
	o := orm.NewOrm()
	ret := Time{Id: id}

	o.Read(&ret)

	return ret
}
