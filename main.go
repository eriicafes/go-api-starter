package main

import (
	"net/http"

	"github.com/eriicafes/go-api-starter/routes"
	"github.com/gin-gonic/gin"
)

// service
type CounterService struct {
	count int
}

func NewCounterService() *CounterService {
	return &CounterService{}
}

func (c *CounterService) Increment() {
	c.count++
}

// controller
type CounterController struct {
	counter *CounterService
}

func NewCounterController(counter *CounterService) *CounterController {
	return &CounterController{
		counter: counter,
	}
}

func (c *CounterController) Register(r *gin.RouterGroup) {
	r.GET("count", c.Count)
}

func (c *CounterController) Count(ctx *gin.Context) {
	c.counter.Increment()

	ctx.JSON(http.StatusOK, gin.H{
		"count": c.counter.count,
	})
}

var (
	counterService    = NewCounterService()
	counterController = NewCounterController(counterService)
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
