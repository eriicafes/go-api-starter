package models

import (
	"errors"
	"time"

	"github.com/eriicafes/filedb"
)

const TodoResource = "todos"

type Todo struct {
	ID          filedb.ID         `json:"id"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	UserID      filedb.ForeignKey `json:"userId"`
	CreatedAt   time.Time         `json:"createdAt"`
	UpdatedAt   time.Time         `json:"updatedAt"`
}

type TodoQuery struct {
	ID     int
	Title  string
	UserID int
}

type todoModel struct {
	db *filedb.Database
}

func NewTodoModel(db *filedb.Database) *todoModel {
	return &todoModel{db}
}

func matchTodoQuery(todo Todo, query *TodoQuery) bool {
	if query == nil {
		return true
	}

	matchID := todo.ID == filedb.ID(query.ID)
	if query.ID == 0 {
		matchID = true
	}

	matchTitle := todo.Title == query.Title
	if query.Title == "" {
		matchTitle = true
	}

	matchUserID := func() bool {
		if todo.UserID != nil {
			return *todo.UserID == filedb.ID(query.UserID)
		}
		return false
	}()
	if query.UserID == 0 {
		matchUserID = true
	}

	return matchID && matchTitle && matchUserID
}

func (m *todoModel) get() []Todo {
	var todos []Todo

	m.db.Get(TodoResource, &todos)

	return todos
}

func (m *todoModel) set(todos []Todo) {
	data := make([]interface{}, len(todos))

	for _, todo := range todos {
		data = append(data, todo)
	}

	m.db.Set(TodoResource, data)
}

func (m *todoModel) FindOneTodo(query *TodoQuery) (*Todo, error) {
	todos := m.get()

	for _, todo := range todos {
		match := matchTodoQuery(todo, query)

		if match {
			return &todo, nil
		}
	}

	return nil, errors.New("todo not found")
}

func (m *todoModel) FindManyTodos(query *TodoQuery) []Todo {
	todos := m.get()
	var result []Todo

	for _, todo := range todos {
		match := matchTodoQuery(todo, query)

		if match {
			result = append(result, todo)
		}
	}

	return result
}

func (m *todoModel) CreateTodo(todo Todo) *Todo {
	todos := m.get()

	// override fields
	todo.ID = filedb.ID(len(todos) + 1)
	todo.CreatedAt = time.Now()
	todo.UpdatedAt = time.Now()

	todos = append(todos, todo)

	m.set(todos)

	return &todo
}

func (m *todoModel) UpdateTodo(query *TodoQuery, updatedTodo Todo) (*Todo, error) {
	todos := m.get()
	var newTodos []Todo

	var updated bool

	for _, todo := range todos {
		if updated {
			newTodos = append(newTodos, todo)
			continue
		}

		match := matchTodoQuery(todo, query)

		if match {
			updated = true

			// override fields
			updatedTodo.ID = todo.ID
			updatedTodo.CreatedAt = todo.CreatedAt
			updatedTodo.UpdatedAt = time.Now()
			if updatedTodo.UserID == nil {
				updatedTodo.UserID = todo.UserID
			}

			newTodos = append(newTodos, updatedTodo)
			continue
		}

		newTodos = append(newTodos, todo)
	}

	if !updated {
		return nil, errors.New("todo not found")
	}

	m.set(newTodos)
	return &updatedTodo, nil
}

func (m *todoModel) RemoveOneTodo(query *TodoQuery) error {
	todos := m.get()
	var newTodos []Todo

	var removed bool

	for _, todo := range todos {
		if removed {
			newTodos = append(newTodos, todo)
			continue
		}

		match := matchTodoQuery(todo, query)

		if match {
			removed = true
			continue
		}

		newTodos = append(newTodos, todo)
	}

	if !removed {
		return errors.New("todo not found")
	}

	m.set(newTodos)
	return nil
}

func (m *todoModel) RemoveManyTodos(query *TodoQuery) (int, error) {
	todos := m.get()
	var newTodos []Todo

	var count int

	for _, todo := range todos {
		match := matchTodoQuery(todo, query)

		if match {
			count++
			continue
		}

		newTodos = append(newTodos, todo)
	}

	if count == 0 {
		return count, errors.New("todos not found")
	}

	m.set(newTodos)

	return count, nil
}
