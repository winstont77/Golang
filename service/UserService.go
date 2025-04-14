package service

type UserService struct{
}

func (userService *UserService) GetUser() string{
	return "test"
}

func (userService *UserService) PostUser(content string) string{
	return content
}