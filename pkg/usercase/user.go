package usercase

import (
	"context"
	"fmt"
	"github.com/THEToilet/events-server/pkg/domain/model"
	"github.com/THEToilet/events-server/pkg/domain/repository"
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

func (u *UserUseCase) GetUser(context.Context) (*model.User, error) {
	user, err := u.userRepository.Find(uuid.New())
	if err != nil {
		fmt.Errorf("unko")
	}
	return user, err
}
