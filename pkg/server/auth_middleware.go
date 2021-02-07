package server

import (
	"github.com/THEToilet/events-server/pkg/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

// AuthMiddleware は認証を担当するミドルウェアを管理する構造体です。
type AuthMiddleware struct {
	authUseCase *usecase.AuthUseCase
}

// NewAuthMiddleware web.AuthMiddlewareのポインタを生成します。
func NewAuthMiddleware(authUseCase *usecase.AuthUseCase) *AuthMiddleware {
	return &AuthMiddleware{
		authUseCase:authUseCase,
	}
}

// Authenticate は認証が必要なAPIで認証情報があるかチェックします。
func (m *AuthMiddleware) Authenticate(nextFunc echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sessCookie, err := c.Cookie("session")
		if err != nil {
			logger.Warnj(map[string]interface{}{"message": "session cookie not found", "error": err.Error()})
			return echo.NewHTTPError(http.StatusUnauthorized)
		}
		userID, err := m.uc.GetUserIDFromSession(sessCookie.Value)
		if err != nil {
			logger.Warnj(map[string]interface{}{"message": "failed to get session", "sessionID": sessCookie.Value, "error": err.Error()})
			return echo.NewHTTPError(http.StatusUnauthorized)
		}

		token, err := m.uc.GetTokenByUserID(userID)
		if err != nil {
			if errors.Is(err, entity.ErrTokenNotFound) {
				logger.Debug(err)
				return echo.NewHTTPError(http.StatusUnauthorized)
			}
			logger.Errorj(map[string]interface{}{"message": "failed to get token", "userID": userID, "sessionID": sessCookie.Value, "error": err.Error()})
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		newToken, err := m.uc.RefreshAccessToken(userID, token)
		if err != nil {
			logger.Errorj(map[string]interface{}{"message": "failed to refresh access token", "sessionID": sessCookie.Value, "error": err.Error()})
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		token = newToken

		ctx := c.Request().Context()
		ctx = service.SetUserIDToContext(ctx, userID)
		ctx = service.SetTokenToContext(ctx, token)
		c.SetRequest(c.Request().WithContext(ctx))
		return nextFunc(c)
	}
}
