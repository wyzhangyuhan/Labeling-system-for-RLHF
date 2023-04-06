package db_model

import (
	"label_system/global/variables"
	"log"
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	*gorm.DB  `gorm:"-" json:"-"`
	CreatedAt time.Time `gorm:"column:creattime" json:"creattime"` //日期时间字段统一设置为字符串即可
	UpdatedAt time.Time `gorm:"column:updatetime" json:"updatetime"`
}

func UseDbConn() *gorm.DB {
	var db *gorm.DB
	if variables.GormDbMysql == nil {
		log.Fatal("no connection!")
	}
	db = variables.GormDbMysql
	return db
}
