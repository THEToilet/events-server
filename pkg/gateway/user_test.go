package gateway

import (
	"errors"
	"github.com/THEToilet/events-server/pkg/domain/model"
	"github.com/THEToilet/events-server/pkg/gateway/database"
	"github.com/google/uuid"
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

func TestUserRepository_FindAll(t *testing.T) {
	// Prepare
	sqlDB, err := database.NewMySqlDB()
	if err != nil {
		t.Fatal(err)
	}
	stmt, err := sqlDB.Prepare("INSERT IGNORE INTO users values('testname1', 'testmail1','testpass1');")
	if err != nil {
		t.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		t.Fatal(err)
	}

	stmt, err = sqlDB.Prepare("INSERT IGNORE INTO users values('testname2', 'testmail2','testpass2');")
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
		want    []*model.User
		wantErr error
	}{
		{
			name: "存在するユーザ一覧を正しく取得できる",
			id:   "testname",
			want: []*model.User{
				{
					UserID:       "testname",
					UserMail:     "testmail",
					UserPassword: "testpass",
				},
				{
					UserID:       "testname1",
					UserMail:     "testmail1",
					UserPassword: "testpass1",
				},
				{
					UserID:       "testname2",
					UserMail:     "testmail2",
					UserPassword: "testpass2",
				},
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &UserRepository{
				sqlDB: sqlDB,
			}
			got, err := r.FindAll()
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("FindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(got, tt.want) {
				t.Errorf("FindAll() diff=%v", cmp.Diff(tt.want, got))
			}
		})
	}

}

func TestUserRepository_Save(t *testing.T) {
	// Prepare
	sqlDB, err := database.NewMySqlDB()
	if err != nil {
		t.Fatal(err)
	}

	testID := uuid.New().String()

	tests := []struct {
		name    string
		id      string
		user    model.User
		want    *model.User
		wantErr error
	}{
		{
			name: "ユーザを正しく登録できる",
			id:   "testname",
			user: model.User{
				UserID:       testID,
				UserMail:     "testmail",
				UserPassword: "testpass",
			},
			want: &model.User{
				UserID:       testID,
				UserMail:     "testmail",
				UserPassword: "testpass",
			},
			wantErr: nil,
		},
		{
			name: "UserIDがかぶった場合エラー",
			id:   "not_found",
			user: model.User{
				UserID:       "testname",
				UserMail:     "testmail",
				UserPassword: "testpass",
			},
			want:    nil,
			wantErr: model.ErrUserAlreadyExisted,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &UserRepository{
				sqlDB: sqlDB,
			}
			err := r.Save(tt.user)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Save() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
	stmt, err := sqlDB.Prepare("DELETE FROM users WHERE user_id = ?;")
	if err != nil {
		t.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(testID)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUserRepository_Delete(t *testing.T) {
	// Prepare
	sqlDB, err := database.NewMySqlDB()
	if err != nil {
		t.Fatal(err)
	}
	testID := uuid.New().String()

	stmt, err := sqlDB.Prepare("INSERT IGNORE INTO users values(?, ?, ?);")
	if err != nil {
		t.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(testID, testID, testID)
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name    string
		id      string
		wantErr bool
	}{
		{
			name:    "存在するユーザを正しく削除できる",
			id:      testID,
			wantErr: false,
		},
		{
			name:    "存在しないUserIDの場合エラ-でない",
			id:      "not_found",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &UserRepository{
				sqlDB: sqlDB,
			}
			if err := r.Delete(tt.id); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
