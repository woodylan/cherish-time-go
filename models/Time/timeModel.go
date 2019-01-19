package TimeModel

import (
	"cherish-time-go/db"
	"time"
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
	CreatedAt  time.Time `gorm:"column:created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at"`
	DeletedAt  time.Time `gorm:"column:deleted_at"`
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

func GetByPage(userId string, perPage, currentPage int) (times []Time, count int, err error) {
	offset := (currentPage - 1) * perPage
	res := db.Conn.Where("user_id = ?", userId).Order("create_time desc").Limit(perPage).Offset(offset).Find(&times).Count(&count)
	err = res.Error

	return
}
