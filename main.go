package main

import (
	_ "label_system/bootstrap"
	"label_system/config"
	"label_system/routes"
	"label_system/scheduler"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	// 0777所有人可读写，0644表示：所有人对可读，创建者可读写
	f, err := os.OpenFile("./mylog.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Panicf("Open log file fail, %v", err)
	}
	log.SetOutput(f)

	if os.Getenv("APPENV") == "local" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	r = routes.SetupRouter(r)
	log.Print("GIN start\n")
	go scheduler.DailyLabelExport()

	r.Run(config.Conf.ServerPort)
}
