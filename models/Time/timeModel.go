package TimeModel

import (
	"github.com/astaxie/beego/orm"
)

// Model Struct
type Time struct {
	Id         string `orm:"column(id);pk"`
	Name       string
	Type       uint8
	Date       string
	Color      string
	Data       string
	Days       string
	Remark     string
	CreateTime string
}

func (a *Time) TableName() string {
	return "time"
}

func GetById(id string) (Time) {
	o := orm.NewOrm()
	u := Time{Id: id}
	//err := o.Read(&u)
	o.Read(&u)

	//fmt.Printf("ERR: %v\n", err)
	//fmt.Println(u.Type)

	return u
}
