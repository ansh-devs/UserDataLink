package database

import (
	"encoding/json"

	"github.com/ansh-devs/tc-assessment/internal/entity"
)

type Repository struct {
	Users []entity.User
}

// GetUsersListByIds
func (repo *Repository) GetUsersListByIds([]int) (users []entity.User) { return users }

// GetUserById
func (repo *Repository) GetUserById(int) (user entity.User) { return user }

// GetUsersByField
func (repo *Repository) GetUsersByField(field, value string) (users []entity.User) { return users }

// NewUserRepository constructor for the UserRepository
func NewUserRepository() UserRepository {
	var users []entity.User

	_ = json.Unmarshal([]byte(MockedUsers), &users)
	// logrus.WithFields(logrus.Fields{"json_data": users}).Info("user_service")
	return &Repository{
		Users: users,
	}
}
