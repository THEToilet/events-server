package usecase

import (
	"context"
	"errors"
	"github.com/THEToilet/events-server/pkg/domain/model"
	"github.com/THEToilet/events-server/pkg/domain/repository"
	"github.com/THEToilet/events-server/pkg/domain/service"
	"github.com/THEToilet/events-server/pkg/log"
	"go.uber.org/zap"
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
	logger := log.New()
	id, ok := service.GetUserIDFromContext(ctx)
	//id := "userID"
	//ok := true
	if ok != true {
		logger.Error("userID not found")
		return nil, errors.New("userID not found")
	}
	if id == "" {
		logger.Error("userID is empty = " + id)
		return nil, errors.New("userID not found")
	}

	user, err := u.userRepository.Find(id)
	if err != nil {
		logger.Error("user not found", zap.Error(err))
		return nil, err
	}
	return user, err
}

func (u *UserUseCase) UserLogin(ctx context.Context) (*model.User, error) {
	return nil, nil
}
