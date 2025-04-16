package service

import(
	"GolangIrisMVC/model"
)

type IUserService interface {
	GetUser() string
	PostUser(user model.User) string
}