package gateway

import (
	"errors"
	"github.com/THEToilet/events-server/pkg/domain/model"
	"github.com/THEToilet/events-server/pkg/gateway/database"
	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
	"testing"
	"time"
)

func TestUserTagRepository_FindAll(t *testing.T) {
	// Prepare
	sqlDB, err := database.NewMySqlDB()
	if err != nil {
		t.Fatal(err)
	}
	testTime := time.Date(2020, 6, 1, 17, 44, 13, 0, time.UTC).String()

	stmt, err := sqlDB.Prepare("INSERT IGNORE INTO users_tags values('testEventID', 'testTagID1', ?, ?);")
	if err != nil {
		t.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(testTime, testTime)
	if err != nil {
		t.Fatal(err)
	}

	stmt, err = sqlDB.Prepare("INSERT IGNORE INTO users_tags values('testEventID', 'testTagID2', ?, ?);")
	if err != nil {
		t.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(testTime, testTime)
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name    string
		id      string
		want    []*model.UserTag
		wantErr error
	}{
		{
			name: "存在するユーザ一タグ覧を正しく取得できる",
			id:   "testname",
			want: []*model.UserTag{
				{
					EventID:   "testEventID",
					TagID:     "testTagName",
					CreatedAt: time.Date(2020, 6, 1, 17, 44, 13, 0, time.UTC),
					UpdatedAt: time.Date(2020, 6, 1, 17, 44, 13, 0, time.UTC),
				},
				{
					EventID:   "testEventID",
					TagID:     "testTagName1",
					CreatedAt: time.Date(2020, 6, 1, 17, 44, 13, 0, time.UTC),
					UpdatedAt: time.Date(2020, 6, 1, 17, 44, 13, 0, time.UTC),
				},
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &UserTagRepository{
				sqlDB: sqlDB,
			}
			got, err := r.FindAll("testEventID")
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

func TestUserTagRepository_Save(t *testing.T) {
	// Prepare
	sqlDB, err := database.NewMySqlDB()
	if err != nil {
		t.Fatal(err)
	}

	testID := uuid.New().String()
	testID1 := uuid.New().String()

	tests := []struct {
		name    string
		id      string
		tag     model.UserTag
		want    *model.UserTag
		wantErr error
	}{
		{
			name: "ユーザを正しく登録できる",
			id:   "testname",
			tag: model.UserTag{
				EventID:   testID,
				TagID:     testID1,
				CreatedAt: time.Date(2020, 6, 1, 17, 44, 13, 0, time.UTC),
				UpdatedAt: time.Date(2020, 6, 1, 17, 44, 13, 0, time.UTC),
			},
			want:    nil,
			wantErr: nil,
		},
		{
			name: "UserIDがかぶった場合エラー",
			id:   "not_found",
			tag: model.UserTag{
				EventID:   testID,
				TagID:     testID1,
				CreatedAt: time.Date(2020, 6, 1, 17, 44, 13, 0, time.UTC),
				UpdatedAt: time.Date(2020, 6, 1, 17, 44, 13, 0, time.UTC),
			},
			want:    nil,
			wantErr: model.ErrUserTagAlreadyExisted,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &UserTagRepository{
				sqlDB: sqlDB,
			}
			err := r.Save(&tt.tag)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Save() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
	stmt, err := sqlDB.Prepare("DELETE FROM users_tags WHERE event_id = ?;")
	if err != nil {
		t.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(testID)
	if err != nil {
		t.Fatal(err)
	}
}

func TestTagRepository_Delete(t *testing.T) {
	// Prepare
	sqlDB, err := database.NewMySqlDB()
	if err != nil {
		t.Fatal(err)
	}
	testID := uuid.New().String()
	testID1 := uuid.New().String()

	stmt, err := sqlDB.Prepare("INSERT IGNORE INTO users_tags values(?, ?, ?, ?);")
	if err != nil {
		t.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(testID, testID1, time.Now(), time.Now())
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
			r := &UserTagRepository{
				sqlDB: sqlDB,
			}
			if err := r.Delete(testID, testID1); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
