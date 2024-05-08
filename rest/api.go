package rest

import (
	"fmt"
	"main/conf"

	"github.com/gin-gonic/gin"
)

func Init() error {

	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	r.GET("/send", sendHandler)
	r.POST("/send", sendHandler)

	r.GET("/check", checkHandler)
	r.POST("/check", checkHandler)

	r.GET("/data", dataHandler)
	r.POST("/data", dataHandler)

	r.GET("/command", commandHandler)
	r.POST("/command", commandHandler)

	r.GET("/alias", aliasHandler)
	r.POST("/alias", aliasHandler)

	r.GET("/restart", restartHandler)
	r.POST("/restart", restartHandler)

	r.GET("/access", accessHandler)
	r.POST("/access", accessHandler)

	r.Run(conf.Config.Server.Host + ":" + fmt.Sprint(conf.Config.Server.Port))

	return nil

}
