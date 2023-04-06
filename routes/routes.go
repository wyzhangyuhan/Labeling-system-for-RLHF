package routes

import (
	"label_system/app/controller"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) *gin.Engine {
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	var backend *gin.RouterGroup
	if os.Getenv("APPENV") == "test" {
		backend = r.Group("/test/labeled_system/")
	} else {
		backend = r.Group("/labeled_system/")
	}

	{
		backend.GET("health", func(ctx *gin.Context) { ctx.String(http.StatusOK, "Hello labeled_backend") })
		userGroup := backend.Group("/users/")
		{
			userGroup.POST("regist", controller.UserRegist)
			userGroup.POST("login", controller.UserLogin)
			userGroup.POST("get_num", controller.GetLabelNumPerP)
		}
		modelGroup := backend.Group("/models/")
		{
			modelGroup.POST("upload", controller.UploadModel)
			modelGroup.POST("delete", controller.DeleteModel)
			modelGroup.GET("get_valid", controller.GetValidModel)
		}
		dataGroup := backend.Group("/data/")
		{
			dataGroup.POST("upload_data", controller.DatasetAdded)
			dataGroup.POST("upload_category", controller.CategoryAdded)
			dataGroup.GET("get_data", controller.GetDataset)
		}
		sessionGroup := backend.Group("/session/")
		{
			// sessionGroup.POST("get_que", controller.RetrieveData)
			sessionGroup.POST("start_session", controller.StartSession)
			sessionGroup.POST("chat", controller.AskModel)
			sessionGroup.POST("infer", controller.PureAsk)
			sessionGroup.POST("edit", controller.SubmitRes)
			sessionGroup.POST("save", controller.Save2OSS)
			sessionGroup.POST("change", controller.ChangeTheRound)
			sessionGroup.POST("more", controller.AskForMore)
		}

	}

	return r
}
