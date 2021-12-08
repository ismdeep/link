package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ismdeep/link/api"
	"github.com/ismdeep/link/config"
)

func main() {
	router := gin.Default()
	router.GET("/", api.EchoHelp)
	router.POST("/api/v1/links", api.LinkAdd)
	router.GET("/:link_id", api.LinkJump)
	router.GET("/404", api.EchoNotFound)
	if err := router.Run(config.Config.Server.Bind); err != nil {
		panic(err)
	}
}
