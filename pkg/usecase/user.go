package usecase

import (
	"context"
	"fmt"
	"github.com/THEToilet/events-server/pkg/domain/model"
	"github.com/THEToilet/events-server/pkg/domain/repository"
	"github.com/THEToilet/events-server/pkg/domain/service"
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

func (u *UserUseCase) GetUser(ctx context.Context) (*model.User, error) {
	id, ok := service.GetUserIDFromContext(ctx)
	if ok != false {
		fmt.Errorf("unko")
	}
	user, err := u.userRepository.Find(id)
	if err != nil {
		fmt.Errorf("unko")
	}
	return user, err
}

func (u *UserUseCase) UserLogin(ctx context.Context) (*model.User, error) {
	return nil, nil
}
