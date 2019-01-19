package UserModel

import (
	"cherish-time-go/db"
	"github.com/jinzhu/gorm"
	"cherish-time-go/modules/util"
	"time"
)

// Model Struct
type User struct {
	Id        string    `gorm:"column(id);pk" json:"id"`
	OpenId    string    `json:"openId"`
	NickName  string    `json:"nickName"`
	Sex       int       `json:"sex"`
	City      string    `json:"city"`
	Province  string    `json:"province"`
	Country   string    `json:"country"`
	Avatar    string    `json:"avatar"`
	CreatedAt time.Time `gorm:"column:created_at" json:"-"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"-"`
	DeletedAt *int64    `gorm:"column:deleted_at" json:"-"`
}

func (a *User) TableName() string {
	return "tb_user"
}

func (user *User) BeforeCreate(scope *gorm.Scope) error {
	user.Id = util.GenShortUuid()
	return nil
}

func GetByOpenId(openId string) (User, error) {
	user := User{}

	res := db.Conn.Where("open_id = ?", openId).First(&user)
	err := res.Error

	return user, err
}

func AddNew(openId, nickName string, sex int, city, province, country, avatar string) (User, bool) {
	user := User{OpenId: openId, NickName: nickName, Sex: sex, City: city, Province: province, Country: country, Avatar: avatar}

	db.Conn.Create(&user)
	res := db.Conn.NewRecord(&user)

	return user, !res
}

func UpdateData(user *User, nickName string, sex int, city, province, country, avatar string) (*User) {
	user.NickName = nickName
	user.Sex = sex
	user.City = city
	user.Province = province
	user.Country = country
	user.Avatar = avatar

	db.Conn.Save(&user)

	return user
}
