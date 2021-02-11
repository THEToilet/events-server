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
	"strings"
	"testing"
	"time"
)

func TestTagHandler_GetTag(t *testing.T) {
	logger := log.New()

	tests := []struct {
		name               string
		userID             string
		prepareMockTagRepo func(repo *mock_repository.MockTagRepository)
		want               *[]tagResponse
		wantErr            bool
		wantCode           int
	}{
		{
			name:   "正しくタグ一覧が取得できる",
			userID: "userID",
			prepareMockTagRepo: func(repo *mock_repository.MockTagRepository) {
				repo.EXPECT().FindAll().Return([]*model.Tag{
					{
						TagID:     "100",
						TagName:   "tag",
						CreatedAt: time.Date(2020, 12, 31, 23, 59, 59, 0, time.Local),
						UpdatedAt: time.Date(2020, 12, 31, 23, 59, 59, 0, time.Local),
					},
					{
						TagID:     "101",
						TagName:   "tag1",
						CreatedAt: time.Date(2020, 12, 31, 23, 59, 59, 0, time.Local),
						UpdatedAt: time.Date(2020, 12, 31, 23, 59, 59, 0, time.Local),
					},
					{
						TagID:     "102",
						TagName:   "tag2",
						CreatedAt: time.Date(2020, 12, 31, 23, 59, 59, 0, time.Local),
						UpdatedAt: time.Date(2020, 12, 31, 23, 59, 59, 0, time.Local),
					},
				}, nil)
			},
			want: &[]tagResponse{
				{
					TagId:     "100",
					TagName:   "tag",
					CreatedAt: time.Date(2020, 12, 31, 23, 59, 59, 0, time.Local),
					UpdatedAt: time.Date(2020, 12, 31, 23, 59, 59, 0, time.Local),
				},
				{
					TagId:     "101",
					TagName:   "tag1",
					CreatedAt: time.Date(2020, 12, 31, 23, 59, 59, 0, time.Local),
					UpdatedAt: time.Date(2020, 12, 31, 23, 59, 59, 0, time.Local),
				},
				{
					TagId:     "102",
					TagName:   "tag2",
					CreatedAt: time.Date(2020, 12, 31, 23, 59, 59, 0, time.Local),
					UpdatedAt: time.Date(2020, 12, 31, 23, 59, 59, 0, time.Local),
				},
			},
			wantErr:  false,
			wantCode: http.StatusOK,
		},
		{
			name:   "タグの取得に失敗したときはInternalServerError",
			userID: "userID",
			prepareMockTagRepo: func(repo *mock_repository.MockTagRepository) {
				repo.EXPECT().FindAll().Return(nil, errors.New("WOO! error"))
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
			repo := mock_repository.NewMockTagRepository(ctrl)
			tt.prepareMockTagRepo(repo)
			tagUseCase := usecase.NewTagUseCase(repo)
			h := &TagHandler{
				tagUseCase: tagUseCase,
			}

			err := h.GetTag(c)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUser() error = %v, wantErr %v", err, tt.wantErr)
			}
			// ステータスコードのチェック
			if er, ok := err.(*echo.HTTPError); (ok && er.Code != tt.wantCode) || (!ok && rec.Code != tt.wantCode) {
				t.Errorf("GetUser() code = %d, want = %d", rec.Code, tt.wantCode)
			}

			if !tt.wantErr {
				got := &[]tagResponse{}
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

func TestTagHandler_PostTag(t *testing.T) {
	logger := log.New()

	tests := []struct {
		name               string
		userID             string
		prepareMockTagRepo func(repo *mock_repository.MockTagRepository)
		wantErr            bool
		wantCode           int
	}{
		{
			name:   "正しくタグが投稿できる",
			userID: "userID",
			prepareMockTagRepo: func(repo *mock_repository.MockTagRepository) {
				repo.EXPECT().Save("tag1").Return(&model.Tag{
					TagID:     "100",
					TagName:   "tag1",
					CreatedAt: time.Date(2020, 12, 31, 23, 59, 59, 0, time.Local),
					UpdatedAt: time.Date(2020, 12, 31, 23, 59, 59, 0, time.Local),
				}, nil)
			},
			wantErr:  false,
			wantCode: http.StatusOK,
		},
		{
			name:   "タグに投稿に失敗したときはInternalServerError",
			userID: "userID",
			prepareMockTagRepo: func(repo *mock_repository.MockTagRepository) {
				repo.EXPECT().Save("tag1").Return(nil, errors.New(" ooo error "))
			},
			wantErr:  true,
			wantCode: http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			jsonStr := `{"tag_name":"` + "tag1" + `"}`
			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(jsonStr))
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
			repo := mock_repository.NewMockTagRepository(ctrl)
			tt.prepareMockTagRepo(repo)
			tagUseCase := usecase.NewTagUseCase(repo)
			h := TagHandler{
				tagUseCase: tagUseCase,
			}

			err := h.PostTag(c)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUser() error = %v, wantErr %v", err, tt.wantErr)
			}
			// ステータスコードのチェック
			if er, ok := err.(*echo.HTTPError); (ok && er.Code != tt.wantCode) || (!ok && rec.Code != tt.wantCode) {
				t.Errorf("GetUser() code = %d, want = %d", rec.Code, tt.wantCode)
			}

		})
	}
}
