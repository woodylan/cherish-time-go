package TimeModel

import (
	"github.com/astaxie/beego/orm"
	"fmt"
	"github.com/astaxie/beego"
)

// Model Struct
type Time struct {
	Id    string `orm:"column(id);pk"`
	Type  int16
	Date  string
	Color string
}

func (a *Time) TableName() string {
	return beego.AppConfig.String("db.prefix") + "time"
}

func GetById(id string) (Time) {
	o := orm.NewOrm()
	u := Time{Id: id}
	err := o.Read(&u)

	fmt.Printf("ERR: %v\n", err)
	fmt.Println(u.Type)

	return u
}
