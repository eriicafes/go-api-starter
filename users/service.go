package users

import (
	"errors"

	"github.com/eriicafes/go-api-starter/models"
)

type UsersService interface {
	FindAll() []models.User
	FindOne(id int) (*models.User, error)
	Create(user models.User) *models.User
	Update(id int, user models.User) (*models.User, error)
	Remove(id int) error
}

type usersService struct {
	usersRepository UsersRepository
}

func NewUsersService(usersRepository UsersRepository) *usersService {
	return &usersService{
		usersRepository: usersRepository,
	}
}

func (s *usersService) FindAll() []models.User {
	return s.usersRepository.FindAll()
}

func (s *usersService) FindOne(id int) (*models.User, error) {
	user, err := s.usersRepository.FindOne(id)

	if err != nil {
		return nil, errors.New("user not found")
	}

	return user, nil
}

func (s *usersService) Create(user models.User) *models.User {
	return s.usersRepository.Create(user)
}

func (s *usersService) Update(id int, user models.User) (*models.User, error) {
	updatedUser, err := s.usersRepository.Update(id, user)

	if err != nil {
		return nil, errors.New("user not found")
	}

	return updatedUser, nil
}

func (s *usersService) Remove(id int) error {
	err := s.usersRepository.Remove(id)

	if err != nil {
		return errors.New("user not found")
	}

	return nil
}
