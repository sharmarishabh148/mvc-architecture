package services

import (
	"github.com/sharmari/mvc/domain"
	"github.com/sharmari/mvc/utils"
)

func GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userId)
}
