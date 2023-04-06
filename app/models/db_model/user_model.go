package db_model

import (
	"log"
)

func CreateUserFactory() *UsersModel {
	return &UsersModel{BaseModel: BaseModel{DB: UseDbConn()}}
}

type UsersModel struct {
	BaseModel
	UserID   int64  `gorm:"column:id;primaryKey"`
	UserName string `gorm:"column:username" json:"user_name"`
	Password string `gorm:"column:password" json:"password"`
}

// 自定义表名
func (u *UsersModel) TableName() string {
	return "users"
}

func (u *UsersModel) Register(username string, psw string, user_id int64) bool {
	res := u.Create(
		&UsersModel{
			UserName: username,
			Password: psw,
			UserID:   user_id,
		})
	if res.Error != nil {
		log.Printf("user数据库插入有误")
		return false

	} else {
		return true
	}
}

func (u *UsersModel) Login(username, psw string) (bool, int64) {
	var um UsersModel
	um.UserName = username

	res := u.Where("username = ?", um.UserName).First(&um)

	if res.Error == nil {
		if psw == um.Password {
			return true, um.UserID
		}
		return false, -1
	} else {
		log.Printf("登录有问题：%v", res.Error)
		return false, -1
	}

}

func (u *UsersModel) GetUserInfo(userid int64) (bool, *UsersModel) {
	var um UsersModel
	um.UserID = userid
	res := u.First(&um)
	if res.Error == nil {
		return true, &um
	}
	return false, nil
}
