package todos

import (
	"net/http"
	"strconv"

	"github.com/eriicafes/go-api-starter/models"
	"github.com/eriicafes/go-api-starter/response"
	"github.com/gin-gonic/gin"
)

type todosController struct {
	todosService TodosService
}

func NewTodosController(todosService TodosService) *todosController {
	return &todosController{
		todosService: todosService,
	}
}

func (c *todosController) Routes(r *gin.RouterGroup) {
	r.GET("", c.FindAll)
	r.GET(":todoId", c.FindOne)
	r.POST("", c.Create)
	r.PUT(":todoId", c.Update)
	r.DELETE(":todoId", c.Remove)
}

func (c *todosController) FindAll(ctx *gin.Context) {
	res := response.New(ctx)
	userId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		res.SetStatusMessage(http.StatusBadRequest, "Invalid ID").ErrJSON()
		return
	}

	todos := c.todosService.FindAll(userId)

	res.SetData(todos).JSON()
}

func (c *todosController) FindOne(ctx *gin.Context) {
	res := response.New(ctx)
	userId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		res.SetStatusMessage(http.StatusBadRequest, "Invalid ID").ErrJSON()
		return
	}

	todoId, err := strconv.Atoi(ctx.Param("todoId"))

	if err != nil {
		res.SetStatusMessage(http.StatusBadRequest, "Invalid ID").ErrJSON()
		return
	}

	todo, err := c.todosService.FindOne(userId, todoId)

	if err != nil {
		res.SetStatusMessage(http.StatusNotFound, err.Error()).ErrJSON()
		return
	}

	res.SetData(todo).JSON()
}

func (c *todosController) Create(ctx *gin.Context) {
	res := response.New(ctx)
	userId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		res.SetStatusMessage(http.StatusBadRequest, "Invalid ID").ErrJSON()
		return
	}

	var todoDto struct {
		Title       string `json:"title" binding:"required"`
		Description string `json:"description" binding:"required"`
	}

	err = ctx.ShouldBindJSON(&todoDto)

	if err != nil {
		res.SetStatusMessage(http.StatusUnprocessableEntity, "validation error").SetError(err.Error()).ErrJSON()
		return
	}

	todo := models.Todo{
		Title:       todoDto.Title,
		Description: todoDto.Description,
	}

	newTodo := c.todosService.Create(userId, todo)

	res.SetData(newTodo).JSON()
}

func (c *todosController) Update(ctx *gin.Context) {
	res := response.New(ctx)
	userId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		res.SetStatusMessage(http.StatusBadRequest, "Invalid ID").ErrJSON()
		return
	}

	todoId, err := strconv.Atoi(ctx.Param("todoId"))

	if err != nil {
		res.SetStatusMessage(http.StatusBadRequest, "Invalid ID").ErrJSON()
		return
	}

	var todoDto struct {
		Title       string `json:"title" binding:"required"`
		Description string `json:"description" binding:"required"`
	}

	err = ctx.ShouldBindJSON(&todoDto)

	if err != nil {
		res.SetStatusMessage(http.StatusUnprocessableEntity, "validation error").SetError(err.Error()).ErrJSON()
		return
	}

	todo := models.Todo{
		Title:       todoDto.Title,
		Description: todoDto.Description,
	}

	newUser, err := c.todosService.Update(userId, todoId, todo)

	if err != nil {
		res.SetStatusMessage(http.StatusNotFound, err.Error()).ErrJSON()
		return
	}

	res.SetData(newUser).JSON()
}

func (c *todosController) Remove(ctx *gin.Context) {
	res := response.New(ctx)
	userId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		res.SetStatusMessage(http.StatusBadRequest, "Invalid ID").ErrJSON()
		return
	}

	todoId, err := strconv.Atoi(ctx.Param("todoId"))

	if err != nil {
		res.SetStatusMessage(http.StatusBadRequest, "Invalid ID").ErrJSON()
		return
	}

	err = c.todosService.Remove(userId, todoId)

	if err != nil {
		res.SetStatusMessage(http.StatusNotFound, err.Error()).ErrJSON()
		return
	}

	res.SetMessage("todo removed").JSON()
}
