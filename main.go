package main

import (
	"net/http"

	"github.com/eriicafes/filedb"
	"github.com/eriicafes/go-api-starter/controller"
	"github.com/eriicafes/go-api-starter/modules/todos"
	"github.com/eriicafes/go-api-starter/modules/users"
	"github.com/gin-gonic/gin"
)

var (
	database        = filedb.New("store")
	usersRepository = users.NewUsersRepository(database)
	usersService    = users.NewUsersService(usersRepository)
	usersController = users.NewUsersController(usersService)
	todosRepository = todos.NewTodosRepository(database)
	todosService    = todos.NewTodosService(todosRepository)
	todosController = todos.NewTodosController(todosService)
)

func main() {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"ping": "pong"})
	})

	binder := controller.NewBinder(router)

	binder.Bind("/users", usersController)
	binder.Bind("/users/:id/todos", todosController)

	router.Run()
}
