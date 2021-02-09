package server

import (
	"github.com/THEToilet/events-server/pkg/domain/service"
	"github.com/THEToilet/events-server/pkg/usecase"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

// AuthMiddleware は認証を担当するミドルウェアを管理する構造体です。
type AuthMiddleware struct {
	authUseCase *usecase.AuthUseCase
}

// NewAuthMiddleware web.AuthMiddlewareのポインタを生成します。
func NewAuthMiddleware(authUseCase *usecase.AuthUseCase) *AuthMiddleware {
	return &AuthMiddleware{
		authUseCase: authUseCase,
	}
}

// Authenticate ユーザ認証を行ってContextへユーザ情報を保存する
func (m *AuthMiddleware) Authenticate(nextFunc echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		// リクエストヘッダからx-token(認証トークン)を取得
		sessCookie, err := c.Cookie("session")
		if err != nil {
			log.Println("x-token is empty")
			return echo.NewHTTPError(http.StatusUnauthorized)
		}

		userID, err := m.authUseCase.GetUserIDFromSession(sessCookie.Value)
		if err != nil {
			log.Println(err)
			return echo.NewHTTPError(http.StatusUnauthorized)
		}

		ctx = service.SetUserIDToContext(ctx, userID)
		c.SetRequest(c.Request().WithContext(ctx))
		return nextFunc(c)
	}
}
