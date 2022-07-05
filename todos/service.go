package todos

import (
	"errors"

	"github.com/eriicafes/go-api-starter/models"
)

type TodosService interface {
	FindAll(userId int) []models.Todo
	FindOne(userId int, id int) (*models.Todo, error)
	Create(userId int, todo models.Todo) *models.Todo
	Update(userId int, id int, todo models.Todo) (*models.Todo, error)
	Remove(userId int, id int) error
}

type todosService struct {
	todosRepository TodosRepository
}

func NewTodosService(todosRepository TodosRepository) *todosService {
	return &todosService{
		todosRepository: todosRepository,
	}
}

func (s *todosService) FindAll(userId int) []models.Todo {
	return s.todosRepository.FindAll(userId)
}

func (s *todosService) FindOne(userId int, id int) (*models.Todo, error) {
	todo, err := s.todosRepository.FindOne(userId, id)

	if err != nil {
		return nil, errors.New("todo not found")
	}

	return todo, nil
}

func (s *todosService) Create(userId int, todo models.Todo) *models.Todo {
	return s.todosRepository.Create(userId, todo)
}

func (s *todosService) Update(userId int, id int, todo models.Todo) (*models.Todo, error) {
	updatedTodo, err := s.todosRepository.Update(userId, id, todo)

	if err != nil {
		return nil, errors.New("todo not found")
	}

	return updatedTodo, nil
}

func (s *todosService) Remove(userId int, id int) error {
	err := s.todosRepository.Remove(userId, id)

	if err != nil {
		return errors.New("todo not found")
	}

	return nil
}
