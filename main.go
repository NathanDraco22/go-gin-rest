package main

import (
	"gin-rest-server-sample/routes"

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
	routes.OnlyPost(router)

	app.Run()

}
