package database

import "github.com/ansh-devs/tc-assessment/internal/entity"

// UserRepository represents a respository that interacts with a database.
type UserRepository interface {
	GetUsersListByIds([]int) (users []entity.User)
	GetUserById(int) (user entity.User)
	GetUsersByField(field, value string) (users []entity.User)
}
