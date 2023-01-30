package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"mimamori-camera-controller/handler"
)

func main() {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://192.168.0.10:8080"}
	config.AllowMethods = []string{"GET"}
	router.Use(cors.New(config))

	api := router.Group("/api/v1")
	{
		api.GET("/exec/change-camera-mode", handler.ExecChangeCameraMode)
		api.GET("/exec/capture-screen", handler.ExecCaptureScreen)
	}
	router.Run(":3000")
}
