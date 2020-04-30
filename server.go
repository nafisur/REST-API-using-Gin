package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nafisur/Gin_HTTP_Framework/controller"
	"github.com/nafisur/Gin_HTTP_Framework/service"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func main() {
	server := gin.Default()

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())

	})

	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.Save(ctx))

	})

	server.Run(":8080")
}
