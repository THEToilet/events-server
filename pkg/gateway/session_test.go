package gateway

import (
	"errors"
	"github.com/THEToilet/events-server/pkg/domain/model"
	"github.com/THEToilet/events-server/pkg/gateway/database"
	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
	"testing"
)

func TestSessionRepository_Find(t *testing.T) {
	// Prepare
	redisConn := database.NewRedis()

	testSessionID := uuid.New().String()
	testUserID := uuid.New().String()

	redisConn.Do("SET", testSessionID, testUserID)

	tests := []struct {
		name    string
		id      string
		want    string
		wantErr error
	}{
		{
			name:    "存在するセッションを正しく取得できる",
			id:      testSessionID,
			want:    testUserID,
			wantErr: nil,
		},
		{
			name:    "存在しないセッションの場合エラー",
			id:      "not_found",
			want:    "",
			wantErr: model.ErrSessionNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &SessionRepository{
				redisConn: redisConn,
			}
			got, err := r.Find(tt.id)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(got, tt.want) {
				t.Errorf("Find() diff=%v", cmp.Diff(tt.want, got))
			}
		})
	}
}


func TestSessionRepository_Save(t *testing.T) {
	// Prepare
	redisConn := database.NewRedis()

	testSessionID := uuid.New().String()
	testUserID := uuid.New().String()

	tests := []struct {
		name      string
		sessionID string
		userID    string
		wantErr   error
	}{
		{
			name:      "存在するセッションを正しく取得できる",
			sessionID: testSessionID,
			userID:    testUserID,
			wantErr:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &SessionRepository{
				redisConn: redisConn,
			}
			err := r.Save(tt.sessionID, tt.userID)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Save() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}

}
