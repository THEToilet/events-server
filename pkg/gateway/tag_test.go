package gateway

import (
	"errors"
	"github.com/THEToilet/events-server/pkg/domain/model"
	"github.com/THEToilet/events-server/pkg/gateway/database"
	"github.com/google/uuid"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestTagRepository_Find(t *testing.T) {
	// Prepare
	sqlDB, err := database.NewMySqlDB()
	if err != nil {
		t.Fatal(err)
	}
	stmt, err := sqlDB.Prepare("INSERT IGNORE INTO tags values('testTagID-123', 'testTagName', ?, ?);")
	if err != nil {
		t.Fatal(err)
	}
	defer stmt.Close()

	testTime := time.Date(2020, 6, 1, 17, 44, 13, 0, time.UTC).String()
	_, err = stmt.Exec(testTime, testTime)
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name    string
		id      string
		want    *model.Tag
		wantErr error
	}{
		{
			name: "存在するタグを正しく取得できる",
			id:   "testTagID-123",
			want: &model.Tag{
				TagID:     "testTagID-123",
				TagName:   "testTagName",
				CreatedAt: time.Date(2020, 6, 1, 17, 44, 13, 0, time.UTC),
				UpdatedAt: time.Date(2020, 6, 1, 17, 44, 13, 0, time.UTC),
			},
			wantErr: nil,
		},
		{
			name:    "存在しないTagIDの場合エラー",
			id:      "not_found",
			want:    nil,
			wantErr: model.ErrTagNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &TagRepository{
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

/*
func TestTagRepository_FindAll(t *testing.T) {
	// Prepare
	sqlDB, err := database.NewMySqlDB()
	if err != nil {
		t.Fatal(err)
	}
	testTime := time.Date(2020, 6, 1, 17, 44, 13, 0, time.Local).String()

	stmt, err := sqlDB.Prepare("INSERT IGNORE INTO tags values('testTagID1', 'testTagName1', ?, ?);")
	if err != nil {
		t.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(testTime, testTime)
	if err != nil {
		t.Fatal(err)
	}

	stmt, err = sqlDB.Prepare("INSERT IGNORE INTO tags values('testTagID2', 'testTagName2', ?, ?);")
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
		want    []*model.Tag
		wantErr error
	}{
		{
			name: "存在するユーザ一覧を正しく取得できる",
			id:   "testname",
			want: []*model.Tag{
				{
					TagID:     "testTagID",
					TagName:   "testTagName",
					CreatedAt: time.Date(2020, 6, 1, 17, 44, 13, 0, time.Local),
					UpdatedAt: time.Date(2020, 6, 1, 17, 44, 13, 0, time.Local),
				},
				{
					TagID:     "testTagID1",
					TagName:   "testTagName1",
					CreatedAt: time.Date(2020, 6, 1, 17, 44, 13, 0, time.Local),
					UpdatedAt: time.Date(2020, 6, 1, 17, 44, 13, 0, time.Local),
				},
				{
					TagID:     "testTagID2",
					TagName:   "testTagName2",
					CreatedAt: time.Date(2020, 6, 1, 17, 44, 13, 0, time.Local),
					UpdatedAt: time.Date(2020, 6, 1, 17, 44, 13, 0, time.Local),
				},
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &TagRepository{
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
*/

func TestTagRepository_Save(t *testing.T) {
	// Prepare
	sqlDB, err := database.NewMySqlDB()
	if err != nil {
		t.Fatal(err)
	}

	testID := uuid.New().String()

	tests := []struct {
		name    string
		id      string
		tag     model.Tag
		want    *model.Tag
		wantErr error
	}{
		{
			name: "ユーザを正しく登録できる",
			id:   "testname",
			tag: model.Tag{
				TagID:     uuid.NewString(),
				TagName:   "testName",
				CreatedAt: time.Date(2020, 6, 1, 17, 44, 13, 0, time.Local),
				UpdatedAt: time.Date(2020, 6, 1, 17, 44, 13, 0, time.Local),
			},
			want: &model.Tag{
				TagID:     uuid.NewString(),
				TagName:   "testName",
				CreatedAt: time.Date(2020, 6, 1, 17, 44, 13, 0, time.Local),
				UpdatedAt: time.Date(2020, 6, 1, 17, 44, 13, 0, time.Local),
			},
			wantErr: nil,
		},
		{
			name: "UserIDがかぶった場合エラー",
			id:   "not_found",
			tag: model.Tag{
				TagID:     "testTagID",
				TagName:   "testTagName",
				CreatedAt: time.Date(2020, 6, 1, 17, 44, 13, 0, time.Local),
				UpdatedAt: time.Date(2020, 6, 1, 17, 44, 13, 0, time.Local),
			},
			want:    nil,
			wantErr: model.ErrTagAlreadyExisted,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &TagRepository{
				sqlDB: sqlDB,
			}
			err := r.Save(&tt.tag)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Save() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
	stmt, err := sqlDB.Prepare("DELETE FROM tags WHERE tag_id = ?;")
	if err != nil {
		t.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(testID)
	if err != nil {
		t.Fatal(err)
	}
}

/*

func TestTagRepository_Delete(t *testing.T) {
	// Prepare
	sqlDB, err := database.NewMySqlDB()
	if err != nil {
		t.Fatal(err)
	}
	testID := uuid.New().String()

	stmt, err := sqlDB.Prepare("INSERT IGNORE INTO tags values(?, ?, ?);")
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

*/
