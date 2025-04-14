package service

type IUserService interface {
	GetUser() string
	PostUser(content string) string 
}