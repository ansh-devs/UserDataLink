package database

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/ansh-devs/tc-assessment/internal/entity"
	"github.com/fatih/structs"
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
	if len(users) == 0 {
		return users, fmt.Errorf("no result found")
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
	lowerCaseField := strings.ToLower(field)
	lowerCaseValue := strings.ToLower(value)
	fmt.Println(lowerCaseField)
	fmt.Println(lowerCaseValue)

	for _, userDAO := range repo.Users {
		m := structs.Map(userDAO)
		for k, v := range m {
			fmt.Println(k+" ", v)
			if strings.ToLower(k) == lowerCaseField && strings.ToLower(fmt.Sprintf("%v", v)) == fmt.Sprintf("%v", lowerCaseValue) {
				users = append(users, userDAO)
			}
		}
	}
	if len(users) == 0 {
		return users, fmt.Errorf("no result found")
	}

	return users, nil
}

// NewUserRepository constructor for the UserRepository
func NewUserRepository() UserRepository {
	var users []entity.User
	_ = json.Unmarshal([]byte(MockedUsers), &users)
	return &Repository{Users: users}
}
