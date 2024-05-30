package database

import "github.com/ansh-devs/tc-assessment/internal/entity"

// UserRepository represents a respository that interacts with a database.
type UserRepository interface {
	GetUsersListByIds([]int64) (users []entity.User, err error)
	GetUserById(int64) (user entity.User, err error)
	GetUsersByField(field, value string) (users []entity.User, err error)
}
