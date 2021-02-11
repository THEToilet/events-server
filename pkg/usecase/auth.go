package usecase

import (
	"github.com/THEToilet/events-server/pkg/domain/repository"
	"github.com/THEToilet/events-server/pkg/log"
	"go.uber.org/zap"
)

type AuthUseCase struct {
	sessionRepository repository.SessionRepository
}

func NewAuthUseCase(sessionRepository repository.SessionRepository) *AuthUseCase {
	return &AuthUseCase{
		sessionRepository: sessionRepository,
	}
}

// GetUserIDFromSession はセッションIDから対応するユーザIDを返します。
func (u *AuthUseCase) GetUserIDFromSession(sessionID string) (string, error) {
	logger := log.New()
	userID, err := u.sessionRepository.Find(sessionID)
	if err != nil {
		logger.Error("get user from session sessionID="+ sessionID, zap.Error(err))
		return "", err
	}
	return userID, nil
}
