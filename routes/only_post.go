package routes

import (
	"gin-rest-server-sample/controllers"

	"github.com/gin-gonic/gin"
)

func OnlyPost(router *gin.RouterGroup) {
	postRouter := router.Group("/post")

	{
		postRouter.POST("/simpleBind", controllers.JsonBindin)
		postRouter.POST("/forcedBind", controllers.JsonForceBindin)
	}

}
