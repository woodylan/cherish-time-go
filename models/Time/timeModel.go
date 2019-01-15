package TimeModel

import (
	"cherish-time-go/db"
)

// Model Struct
type Time struct {
	Id         string `gorm:"column(id);pk"`
	Name       string
	UserId     string
	Type       uint8
	Date       string
	Color      string
	Remark     string

	CreateTime int64
	CreatedAt int64 `gorm:"column(create_time)"`
	UpdatedAt int64 `gorm:"column(update_time)"`
	DeletedAt *int64
}

func (a *Time) TableName() string {
	return "tb_time"
}

func GetById(id string) (Time, error) {
	ret := Time{Id: id}

	res := db.Conn.Take(&ret)
	err := res.Error

	return ret, err
}
