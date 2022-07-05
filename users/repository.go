package users

import (
	"github.com/eriicafes/filedb"
	"github.com/eriicafes/go-api-starter/models"
)

type UsersRepository interface {
	FindAll() []models.User
	FindOne(id int) (*models.User, error)
	Create(user models.User) *models.User
	Update(id int, user models.User) (*models.User, error)
	Remove(id int) error
}

type usersRepository struct {
	db *filedb.Database
}

func NewUsersRepository(db *filedb.Database) *usersRepository {
	return &usersRepository{
		db: db,
	}
}

func (r *usersRepository) FindAll() []models.User {
	model := models.NewUserModel(r.db)

	return model.FindManyUsers(nil)
}

func (r *usersRepository) FindOne(id int) (*models.User, error) {
	model := models.NewUserModel(r.db)

	return model.FindOneUser(&models.UserQuery{
		ID: id,
	})
}

func (r *usersRepository) Create(user models.User) *models.User {
	model := models.NewUserModel(r.db)

	return model.CreateUser(user)
}

func (r *usersRepository) Update(id int, user models.User) (*models.User, error) {
	model := models.NewUserModel(r.db)

	return model.UpdateUser(&models.UserQuery{
		ID: id,
	}, user)
}

func (r *usersRepository) Remove(id int) error {
	model := models.NewUserModel(r.db)

	return model.RemoveOneUser(&models.UserQuery{
		ID: id,
	})
}
