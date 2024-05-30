package database

import (
	"encoding/json"

	"github.com/ansh-devs/tc-assessment/internal/entity"
)

type Repository struct {
	Users []entity.User
}

// GetUsersListByIds
func (repo *Repository) GetUsersListByIds(ids []int64) (users []entity.User, err error) {
	for _, id := range ids {
		for _, user_iter := range repo.Users {
			if id == user_iter.ID {
				users = append(users, user_iter)
			}
		}
	}
	return users, nil
}

// GetUserById
func (repo *Repository) GetUserById(userID int64) (user entity.User, err error) {
	for _, v := range repo.Users {
		if v.ID == userID {
			user = v
			return
		}
	}
	return user, nil
}

// GetUsersByField
func (repo *Repository) GetUsersByField(field, value string) (users []entity.User, err error) {
	return repo.Users, nil
}

// NewUserRepository constructor for the UserRepository
func NewUserRepository() UserRepository {
	var users []entity.User
	_ = json.Unmarshal([]byte(MockedUsers), &users)
	return &Repository{Users: users}
}
