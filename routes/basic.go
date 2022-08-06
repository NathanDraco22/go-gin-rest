package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BasicsRoutes(router *gin.RouterGroup) {
	basicsRoutes := router.Group("/basic")

	{
		basicsRoutes.GET("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"ok":  true,
				"msg": "json crated by gin.H serializer",
			})
		})
	}

}
