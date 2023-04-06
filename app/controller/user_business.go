package controller

import (
	"fmt"
	"label_system/app/curd/users"
	"label_system/app/models/req_model"
	"label_system/app/response"
	"label_system/global/consts"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UserLogin(c *gin.Context) {
	var userInfo req_model.LoginReq
	if err := c.ShouldBindJSON(&userInfo); err != nil {
		response.Fail(c, consts.BackendError, "", err)
	}
	//v1先不加password统一转成123
	userInfo.Password = "123"
	ok, userId := users.CreateUserCurdFactory().Login(userInfo.Username, userInfo.Password)
	if !ok {
		response.Fail(c, consts.BackendError, "", "登录失败，数据库异常")
		return
	}
	if userId == -1 {
		response.Fail(c, consts.CurdUserError, "", "登陆失败, 用户不存在")
		return
	}
	fmt.Printf("userid: %v", userId)
	response.Success(c, "", map[string]string{"userid": strconv.FormatInt(userId, 10)})

}

func UserRegist(c *gin.Context) {
	var userInfo req_model.LoginReq
	if err := c.ShouldBindJSON(&userInfo); err != nil {
		response.Fail(c, consts.BackendError, "", err)
	}
	//v1先不加password统一转成123'
	userInfo.Password = "123"
	ok, userId := users.CreateUserCurdFactory().Register(userInfo.Username, userInfo.Password)
	if !ok {
		response.Fail(c, consts.CurdUserError, "", "")
		return
	}
	response.Success(c, "", map[string]string{"userid": strconv.FormatInt(userId, 10)})
}
