package main

import (
	"fmt"
	"gin-rest-server-sample/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AnyReponse struct {
	Ok  bool   `json:"ok"`
	Msg string `json:"msg"`
}

func main() {

	app := gin.Default()

	router := app.Group("/api")

	routes.BasicsRoutes(router)

	app.GET("/", func(ctx *gin.Context) {
		fmt.Println(ctx.Request.Header)
		ctx.JSON(http.StatusOK, AnyReponse{
			Ok:  true,
			Msg: "Hello World",
		})
	})

	app.Run()

}
