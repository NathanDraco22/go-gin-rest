package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BasicController(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"ok":  true,
		"msg": "from controller in basic route",
	})
}
