package gateway

import (
	"errors"
	"github.com/THEToilet/events-server/pkg/domain/model"
	"github.com/THEToilet/events-server/pkg/gateway/database"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestUserRepository_Find(t *testing.T) {
	// Prepare
	sqlDB, err := database.NewMySqlDB()
	if err != nil {
		t.Fatal(err)
	}
	stmt, err := sqlDB.Prepare("INSERT IGNORE INTO users values('testname', 'testmail','testpass');")
	if err != nil {
		t.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name    string
		id      string
		want    *model.User
		wantErr error
	}{
		{
			name: "存在するユーザを正しく取得できる",
			id:   "testname",
			want: &model.User{
				UserID:       "testname",
				UserMail:     "testmail",
				UserPassword: "testpass",
			},
			wantErr: nil,
		},
		{
			name:    "存在しないUserIDの場合エラー",
			id:      "not_found",
			want:    nil,
			wantErr: model.ErrUserNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &UserRepository{
				sqlDB: sqlDB,
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
