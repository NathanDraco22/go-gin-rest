package routes

import (
	"gin-rest-server-sample/controllers"

	"github.com/gin-gonic/gin"
)

func BasicsRoutes(router *gin.RouterGroup) {
	basicsRoutes := router.Group("/basic")

	{
		basicsRoutes.GET("/", controllers.BasicController)
		basicsRoutes.GET("/struct", controllers.BasicStructResponse)
	}

}
