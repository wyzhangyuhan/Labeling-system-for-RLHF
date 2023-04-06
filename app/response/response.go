package response

import (
	"fmt"
	"label_system/global/consts"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReturnJson(Context *gin.Context, httpCode int, dataCode int, msg string, data interface{}) {

	//Context.Header("key2020","value2020")  	//可以根据实际情况在头部添加额外的其他信息
	Context.JSON(httpCode, gin.H{
		"code": dataCode,
		"msg":  msg,
		"data": data,
	})
}

// Success 直接返回成功
func Success(c *gin.Context, msg string, data interface{}) {
	ReturnJson(c, http.StatusOK, consts.BackendSuccess, msg, data)
}

func Fail(c *gin.Context, dataCode int, msg string, data interface{}) {
	fmt.Printf("err data: %v", data)
	ReturnJson(c, http.StatusBadRequest, dataCode, msg, data)
	c.Abort()
}
