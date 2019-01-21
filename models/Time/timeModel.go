package TimeModel

import (
	"cherish-time-go/db"
	"time"
	"github.com/jinzhu/gorm"
	"cherish-time-go/modules/util"
)

// Model Struct
type Time struct {
	Id           string `gorm:"column(id);pk"`
	Name         string
	UserId       string
	Type         uint8
	Date         string
	Color        string
	Remark       string
	CreateUserId string
	UpdateUserId string
	CreatedAt    time.Time `gorm:"column:created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at"`
	DeletedAt    time.Time `gorm:"column:deleted_at"`
}

func (time *Time) BeforeCreate(scope *gorm.Scope) error {
	time.Id = util.GenShortUuid()
	return nil
}

func (a *Time) TableName() string {
	return "tb_time"
}

func AddNew(name, userId string, TimeType uint8, date, color, remark string) (Time, bool) {
	time := Time{Name: name, UserId: userId, Type: TimeType, Date: date, Color: color, Remark: remark, CreateUserId: userId, UpdateUserId: userId}

	db.Conn.Create(&time)
	res := db.Conn.NewRecord(&time)

	return time, !res
}

func GetById(id string) (Time, error) {
	ret := Time{Id: id}

	res := db.Conn.Where("id = ?", id).Find(&ret)
	err := res.Error

	return ret, err
}

func GetByPage(userId string, perPage, currentPage int) (times []Time, count int, err error) {
	offset := (currentPage - 1) * perPage
	res := db.Conn.Where("user_id = ?", userId).Order("create_at desc").Limit(perPage).Offset(offset).Find(&times).Count(&count)
	err = res.Error

	return
}
