package bootstrap

import (
	"fmt"
	"label_system/config"
	"label_system/global/variables"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {

	// 6.根据配置初始化 gorm mysql 全局 *gorm.Db
	username := config.Conf.SqlUsername
	password := config.Conf.SqlPassword
	host := config.Conf.SqlHost
	db_name := config.Conf.SqlName
	// SetMaxIdleConns := 10
	// SetMaxOpenConns := 128
	// SetConnMaxLifetime := 60

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True", username, password, host, db_name)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn}), &gorm.Config{})
	if err != nil {
		log.Panic("连接数据库失败")
	}
	variables.GormDbMysql = db
}
