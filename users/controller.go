package users

import (
	"net/http"
	"strconv"

	"github.com/eriicafes/go-api-starter/models"
	"github.com/eriicafes/go-api-starter/response"
	"github.com/gin-gonic/gin"
)

type usersController struct {
	usersService UsersService
}

func NewUsersController(usersService UsersService) *usersController {
	return &usersController{
		usersService: usersService,
	}
}

func (c *usersController) Register(r *gin.RouterGroup) {
	r.GET("", c.FindAll)
	r.GET(":id", c.FindOne)
	r.POST("", c.Create)
	r.PUT(":id", c.Update)
	r.DELETE(":id", c.Remove)
}

func (c *usersController) FindAll(ctx *gin.Context) {
	res := response.New(ctx)

	users := c.usersService.FindAll()

	res.SetData(users).JSON()
}

func (c *usersController) FindOne(ctx *gin.Context) {
	res := response.New(ctx)
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		res.SetStatusMessage(http.StatusBadRequest, "Invalid ID").ErrJSON()
		return
	}

	user, err := c.usersService.FindOne(id)

	if err != nil {
		res.SetStatusMessage(http.StatusNotFound, err.Error()).ErrJSON()
		return
	}

	res.SetData(user).JSON()
}

func (c *usersController) Create(ctx *gin.Context) {
	res := response.New(ctx)

	var userDto struct {
		Name string `json:"name" binding:"required"`
		Age  int    `json:"age" binding:"required"`
	}

	err := ctx.ShouldBindJSON(&userDto)

	if err != nil {
		res.SetStatusMessage(http.StatusUnprocessableEntity, "validation error").SetError(err.Error()).ErrJSON()
		return
	}

	user := models.User{
		Name: userDto.Name,
		Age:  userDto.Age,
	}

	newUser := c.usersService.Create(user)

	res.SetData(newUser).JSON()
}

func (c *usersController) Update(ctx *gin.Context) {
	res := response.New(ctx)
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		res.SetStatusMessage(http.StatusBadRequest, "Invalid ID").ErrJSON()
		return
	}

	var userDto struct {
		Name string `json:"name" binding:"required"`
		Age  int    `json:"age" binding:"required"`
	}

	err = ctx.ShouldBindJSON(&userDto)

	if err != nil {
		res.SetStatusMessage(http.StatusUnprocessableEntity, "validation error").SetError(err.Error()).ErrJSON()
		return
	}

	user := models.User{
		Name: userDto.Name,
		Age:  userDto.Age,
	}

	newUser, err := c.usersService.Update(id, user)

	if err != nil {
		res.SetStatusMessage(http.StatusNotFound, err.Error()).ErrJSON()
		return
	}

	res.SetData(newUser).JSON()
}

func (c *usersController) Remove(ctx *gin.Context) {
	res := response.New(ctx)
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		res.SetStatusMessage(http.StatusBadRequest, "Invalid ID").ErrJSON()
		return
	}

	err = c.usersService.Remove(id)

	if err != nil {
		res.SetStatusMessage(http.StatusNotFound, err.Error()).ErrJSON()
		return
	}

	res.SetMessage("user removed").JSON()
}
