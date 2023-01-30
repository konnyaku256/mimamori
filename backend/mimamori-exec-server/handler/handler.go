package handler

import (
	"github.com/gin-gonic/gin"

	"mimamori-exec-server/exec_command"
)

func ExecChangeCameraMode(c *gin.Context) {
	cameraMode := c.Query("mode")
	err := exec_command.ExecCommand(exec_command.SelectCommandByCameraMode(cameraMode))
	MakeResponse(c, err)
}

func ExecCaptureScreen(c *gin.Context) {
	err := exec_command.ExecCommand(exec_command.CaptureScreen)
	MakeResponse(c, err)	
}
