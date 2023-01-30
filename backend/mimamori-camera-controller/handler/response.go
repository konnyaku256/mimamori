package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func MakeResponse(c *gin.Context, err error) {
	if err != nil {
		c.Status(http.StatusInternalServerError)
	} else {
		c.Status(http.StatusOK)
	}
}