package main

import (
	"fmt"
	"net/http"
	"os/exec"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"mimamori-camera-controller/v4l2"
)

func main() {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://192.168.0.10:8080"}
	config.AllowMethods = []string{"GET"}
	router.Use(cors.New(config))

	api := router.Group("/api/v1")
	{
		api.GET("/exec-v4l2", execV4l2)
	}
	router.Run(":3000")
}

func execV4l2(c *gin.Context) {
	mode := c.Query("mode")

	cc := v4l2.SelectCameraControl(mode)
	out, err := exec.Command(cc.CommandName, cc.CommandArguments...).CombinedOutput()
	fmt.Printf("Command result:\n%s :Error:\n%v\n", out, err)
	if err != nil {
		c.Status(http.StatusBadRequest)
	} else {
		c.Status(http.StatusOK)
	}
}