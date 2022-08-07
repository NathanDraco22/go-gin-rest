package routes

import (
	"fmt"
	"gin-rest-server-sample/models"

	"github.com/gin-gonic/gin"
)

func OnlyPost(router *gin.RouterGroup) {
	postRouter := router.Group("/post")

	{
		postRouter.POST("/", func(ctx *gin.Context) {
			// Declaring a default value Struct
			var anyBody models.RequestBody
			// This can auto serial a json body in a Struct
			ctx.BindJSON(&anyBody)

			fmt.Println(anyBody)
		})
	}

}
