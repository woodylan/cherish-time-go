package TimeModel

import (
	"cherish-time-go/db"
	"fmt"
)

//var db *gorm.DB

// Model Struct
type Time struct {
	//gorm.Model
	Id         string `gorm:"column(id);pk"`
	Name       string
	UserId     string
	Type       uint8
	Date       string
	Color      string
	Remark     string
	CreateTime int64
}

func (a *Time) TableName() string {
	return "tb_time"
}

func GetById(id string) (Time) {
	ret := Time{Id: id}

	db.Conn.Take(&ret)
	fmt.Println(ret)

	return ret
}
