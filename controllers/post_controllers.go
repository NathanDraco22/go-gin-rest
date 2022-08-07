package controllers

import (
	"fmt"
	"gin-rest-server-sample/models"

	"github.com/gin-gonic/gin"
)

func JsonBindin(ctx *gin.Context) {
	var body models.RequestBody
	// bind json without validation
	ctx.BindJSON(&body)

	fmt.Println(body)
}

func JsonForceBindin(ctx *gin.Context) {
	var body models.RequestBody
	// force and validating json, must provide a body in request
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(400, gin.H{
			"error": "this is a error",
		})
		return
	}
	if body.Msg == "" {
		ctx.JSON(400, gin.H{
			"msg": "invalid body",
		})
		return
	}
	ctx.JSON(200, gin.H{"msg": "nice request"})

}
