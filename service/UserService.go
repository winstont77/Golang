package service

import(
	"GolangIrisMVC/model"
	"GolangIrisMVC/infrastructure"
	"log"
)

type UserService struct{
}

func (userService *UserService) GetUser() string{
	return "test"
}

func (userService *UserService) PostUser(user model.User) string{
	db := infrastructure.DB

	if err := db.Create(&user).Error; err != nil {
		log.Println("Error creating user:", err)
		return "failed"
	}

	return "success"
}