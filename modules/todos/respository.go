package todos

import (
	"github.com/eriicafes/filedb"
	"github.com/eriicafes/go-api-starter/models"
)

type TodosRepository interface {
	FindAll(userId int) []models.Todo
	FindOne(userId int, id int) (*models.Todo, error)
	Create(userId int, todo models.Todo) *models.Todo
	Update(userId int, id int, todo models.Todo) (*models.Todo, error)
	Remove(userId int, id int) error
}

type todosRepository struct {
	db *filedb.Database
}

func NewTodosRepository(db *filedb.Database) *todosRepository {
	return &todosRepository{
		db: db,
	}
}

func (r *todosRepository) FindAll(userId int) []models.Todo {
	model := models.NewTodoModel(r.db)

	return model.FindManyTodos(&models.TodoQuery{
		UserID: userId,
	})
}

func (r *todosRepository) FindOne(userId int, id int) (*models.Todo, error) {
	model := models.NewTodoModel(r.db)

	return model.FindOneTodo(&models.TodoQuery{
		UserID: userId,
		ID:     id,
	})
}

func (r *todosRepository) Create(userId int, todo models.Todo) *models.Todo {
	model := models.NewTodoModel(r.db)

	foreignId := filedb.ID(userId)
	todo.UserID = &foreignId

	return model.CreateTodo(todo)
}

func (r *todosRepository) Update(userId int, id int, todo models.Todo) (*models.Todo, error) {
	model := models.NewTodoModel(r.db)

	return model.UpdateTodo(&models.TodoQuery{
		UserID: userId,
		ID:     id,
	}, todo)
}

func (r *todosRepository) Remove(userId int, id int) error {
	model := models.NewTodoModel(r.db)

	return model.RemoveOneTodo(&models.TodoQuery{
		UserID: userId,
		ID:     id,
	})
}
