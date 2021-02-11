package server

import (
	"errors"
	"github.com/THEToilet/events-server/pkg/domain/mock_repository"
	"github.com/THEToilet/events-server/pkg/usecase"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuthMiddleware_Authenticate(t *testing.T) {
	tests := []struct {
		name                     string
		prepareRequest           func(req *http.Request)
		prepareSessionRepository func(r *mock_repository.MockSessionRepository)
		next                     echo.HandlerFunc
		wantErr                  bool
		wantCode                 int
	}{
		{
			name: "セッションがクッキーに存在しないと401",
			prepareRequest: func(req *http.Request) {
			},
			prepareSessionRepository: func(r *mock_repository.MockSessionRepository) {},
			next:                     nil,
			wantErr:                  true,
			wantCode:                 http.StatusUnauthorized,
		},
		{
			name: "DBからセッションの取得に成功",
			prepareRequest: func(req *http.Request) {
				req.AddCookie(&http.Cookie{
					Name:   "session",
					Value:  "sessionID",
					Path:   "/",
					MaxAge: 60 * 60 * 24 * 7,
				})
			},
			prepareSessionRepository: func(r *mock_repository.MockSessionRepository) {
				r.EXPECT().Find("sessionID").Return("11111", nil)
			},
			next:     nil,
			wantErr:  true,
			wantCode: http.StatusUnauthorized,
		},
		{
			name: "DBからセッションの取得に失敗すると401",
			prepareRequest: func(req *http.Request) {
				req.AddCookie(&http.Cookie{
					Name:   "session",
					Value:  "sessionID",
					Path:   "/",
					MaxAge: 60 * 60 * 24 * 7,
				})
			},
			prepareSessionRepository: func(r *mock_repository.MockSessionRepository) {
				r.EXPECT().Find("sessionID").Return("", errors.New("unknown error"))
			},
			next:     nil,
			wantErr:  true,
			wantCode: http.StatusUnauthorized,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// httptestの準備
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			tt.prepareRequest(req)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			// モックの準備
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			sessionRepository := mock_repository.NewMockSessionRepository(ctrl)
			tt.prepareSessionRepository(sessionRepository)

			m := &AuthMiddleware{
				authUseCase: usecase.NewAuthUseCase(sessionRepository),
			}
			err := m.Authenticate(tt.next)(c)
			if (err != nil) != tt.wantErr {
				t.Errorf("AuthMiddleware.Authenticate() error = %v, wantErr %v", err, tt.wantErr)
			}

			// ステータスコードのチェック
			if er, ok := err.(*echo.HTTPError); (ok && er.Code != tt.wantCode) || (!ok && rec.Code != tt.wantCode) {
				t.Errorf("AuthMiddleware.Authenticate() code = %d, want = %d", rec.Code, tt.wantCode)
			}
		})
	}
}
