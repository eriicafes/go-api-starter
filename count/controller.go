package count

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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
