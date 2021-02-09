package usecase

import (
	"fmt"
	"github.com/THEToilet/events-server/pkg/domain/repository"
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
	userID, err := u.sessionRepository.Find(sessionID)
	if err != nil {
		return "", fmt.Errorf("get user from session sessionID=%s: %w", sessionID, err)
	}
	return userID, nil
}
