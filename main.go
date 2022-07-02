package main

import (
	"net/http"

	"github.com/eriicafes/go-api-starter/count"
	"github.com/eriicafes/go-api-starter/routes"
	"github.com/gin-gonic/gin"
)

var (
	counterService    = count.NewCounterService()
	counterController = count.NewCounterController(counterService)
)

func main() {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"ping": "pong"})
	})

	routes.RegisterControllers(router,
		routes.Register("/", counterController),
		routes.Register("/second", counterController),
	)

	router.Run()
}
