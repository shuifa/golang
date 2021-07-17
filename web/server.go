package main

import (
	"github.com/gin-gonic/gin"
	"githun.com/oushuifa/golang/web/controllers"
	"log"
	"net/http"
)

func main() {

	server := gin.Default()

	server.GET("/ping", func(context *gin.Context) {
		context.String(http.StatusOK, "%s", "pong")
	})

	videoCtrl := controllers.New()

	server.GET("/videos", videoCtrl.GetAll)
	server.PUT("/videos/:id", videoCtrl.Update)
	server.POST("/videos", videoCtrl.Create)
	server.DELETE("/videos/:id", videoCtrl.Delete)

	log.Fatalln(server.Run(":8080"))
}

