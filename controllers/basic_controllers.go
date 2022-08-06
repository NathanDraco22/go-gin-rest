package controllers

import (
	"gin-rest-server-sample/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BasicController(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"ok":  true,
		"msg": "from controller in basic route",
	})
}

func BasicStructResponse(ctx *gin.Context) {

	response := &models.BasicResponse{
		Ok:        true,
		Msg:       "Correct Request",
		ExtraInfo: "You fetched from memory * struct",
	}

	ctx.JSON(http.StatusOK, response)
}
