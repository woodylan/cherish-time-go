package SentenceModel

import (
	"cherish-time-go/db"
	"time"
	"github.com/jinzhu/gorm"
)

// Model Struct
type Sentence struct {
	Id           string    `gorm:"column(id);pk" json:"id"`
	Content      string    `json:"content"`
	Author       string    `json:"author"`
	Book         string    `json:"book"`
	ShowTimes    uint16    `json:"-"`
	CreateUserId string    `json:"-"`
	UpdateUserId string    `json:"-"`
	CreatedAt    time.Time `gorm:"column:created_at" json:"-"`
	UpdatedAt    time.Time `gorm:"column:updated_at" json:"-"`
	DeletedAt    *time.Time `gorm:"column:deleted_at" json:"-"`
}

func (a *Sentence) TableName() string {
	return "tb_sentence"
}

func GetRand(perPage int) (sentences []Sentence, err error) {
	//db.Conn.LogMode(true)
	res := db.Conn.Order("RAND()").Limit(perPage).Find(&sentences)
	err = res.Error
	if err == nil {
		keyList := make([]string, 0)
		for _, val := range sentences {
			keyList = append(keyList, val.Id)
		}

		//自增1
		db.Conn.Table("tb_sentence").Where("id in (?)", keyList).Update("show_times", gorm.Expr("show_times + 1"))
	}

	return
}
