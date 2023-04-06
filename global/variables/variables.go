package variables

import "gorm.io/gorm"

var (
	GormDbMysql *gorm.DB // 全局gorm的客户端连接
)
