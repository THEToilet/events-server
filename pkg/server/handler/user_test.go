package handler

import (
	"encoding/json"
	"errors"
	"github.com/THEToilet/events-server/pkg/domain/mock_repository"
	"github.com/THEToilet/events-server/pkg/domain/model"
	"github.com/THEToilet/events-server/pkg/domain/service"
	"github.com/THEToilet/events-server/pkg/log"
	"github.com/THEToilet/events-server/pkg/usecase"
	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/labstack/echo/v4"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserHandler_GetUser(t *testing.T) {
	//t.Parallel()
	logger := log.New()

	tests := []struct {
		name                string
		userID              string
		prepareMockUserRepo func(repo *mock_repository.MockUserRepository)
		want                *userResponse
		wantErr             bool
		wantCode            int
	}{
		{
			name:   "正しくメルアドが取得できる",
			userID: "userID",
			prepareMockUserRepo: func(repo *mock_repository.MockUserRepository) {
				repo.EXPECT().Find("userID").Return(&model.User{
					UserID:   "userID",
					UserMail: "test@mail.com",
				}, nil)
			},
			want: &userResponse{
				Mail: "test@mail.com",
			},
			wantErr:  false,
			wantCode: http.StatusOK,
		},
		{
			name:   "DBからユーザの取得に失敗したときはInternalServerError",
			userID: "userID",
			prepareMockUserRepo: func(repo *mock_repository.MockUserRepository) {
				repo.EXPECT().Find("userID").Return(nil, errors.New("WOO! error"))
			},
			want:     nil,
			wantErr:  true,
			wantCode: http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			ctx := c.Request().Context()
			ctx = service.SetUserIDToContext(ctx, tt.userID)
			c.SetRequest(c.Request().WithContext(ctx))
			val, _ := service.GetUserIDFromContext(c.Request().Context())
			logger.Info(val)

			// モックの準備
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			repo := mock_repository.NewMockUserRepository(ctrl)
			tt.prepareMockUserRepo(repo)
			userUseCase := usecase.NewUserUseCase(repo)
			h := &UserHandler{
				userUseCase: userUseCase,
			}

			err := h.GetUser(c)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUser() error = %v, wantErr %v", err, tt.wantErr)
			}
			// ステータスコードのチェック
			if er, ok := err.(*echo.HTTPError); (ok && er.Code != tt.wantCode) || (!ok && rec.Code != tt.wantCode) {
				t.Errorf("GetUser() code = %d, want = %d", rec.Code, tt.wantCode)
			}

			if !tt.wantErr {
				got := &userResponse{}
				err := json.Unmarshal(rec.Body.Bytes(), got)
				if err != nil {
					t.Fatal(err)
				}
				if !cmp.Equal(got, tt.want) {
					t.Errorf("GetUser() diff = %v", cmp.Diff(got, tt.want))
				}
			}
		})
	}
}
