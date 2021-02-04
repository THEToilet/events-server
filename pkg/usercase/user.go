package usercase

import (
	"../domain/repository"
	"../domain/model"
	"fmt"
	"github.com/google/uuid"
)

//UserUseCase ユーザへのアプリケーション層です
type UserUseCase struct {
	userRepository repository.UserRepository
}

//NewUserUseCase
func NewUserUseCase(userRepository repository.UserRepository) *UserUseCase {
	return &UserUseCase{
		userRepository: userRepository,
	}
}

func (u *UserUseCase) GetUser() (*model.User, error){
	user, err := u.userRepository.Find(uuid.New())
	if err != nil {
		fmt.Errorf("unko")
	}
	return user, err
}
