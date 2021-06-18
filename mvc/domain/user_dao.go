package domain

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sharmari/mvc/utils"
)

var (
	users = map[int64]*User{
		123: {Id: 1, FirstName: "Rishabh",
			LastName: "Sharma", Email: "sharmarishabh148@gmail.com"},
	}
	UserDao userDaoInterface
)

func init() {
	UserDao = &userDao{}
}

type userDaoInterface interface {
	GetUser(int64) (*User, *utils.ApplicationError)
}

type userDao struct{}

func (u *userDao) GetUser(userId int64) (*User, *utils.ApplicationError) {
	log.Println("we're accessing the database")
	if user := users[userId]; user != nil {
		return user, nil
	}

	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user %v does not exists", userId),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}
