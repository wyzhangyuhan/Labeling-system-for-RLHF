package users

import (
	"label_system/app/models/db_model"
	"label_system/utils/md5_encrypt"
	"label_system/utils/snowflakes"
)

func CreateUserCurdFactory() *UserCurd {
	return &UserCurd{db_model.CreateUserFactory()}
}

type UserCurd struct {
	userModel *db_model.UsersModel
}

func (u *UserCurd) Register(username, psw string) (bool, int64) {
	psw = md5_encrypt.B64MD5(psw)
	userid := snowflakes.CreateSnowflakeFactory().GenId()
	return u.userModel.Register(username, psw, userid), userid
}

func (u *UserCurd) Login(username, psw string) (bool, int64) {
	psw = md5_encrypt.B64MD5(psw)
	return u.userModel.Login(username, psw)
}

func (u *UserCurd) GetUserInfo(userid int64) (bool, *db_model.UsersModel) {

	return u.userModel.GetUserInfo(userid)
}
