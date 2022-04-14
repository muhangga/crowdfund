package service

import (
	"github.com/muhangga/entity"
	model "github.com/muhangga/model/request"
)

type UserService interface {
	RegisterUser(userRequest model.RegisterRequest) (entity.User, error)
}