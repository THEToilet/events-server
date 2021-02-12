package gateway

import (
	"database/sql"
	"github.com/THEToilet/events-server/pkg/domain/model"
	"github.com/THEToilet/events-server/pkg/domain/repository"
	"github.com/google/uuid"
	"time"
)

var _ repository.TagRepository = &TagRepository{}

// UserRepository は repository.TagRepository を満たす構造体です
type TagRepository struct {
	sqlDB *sql.DB
}

// NewUserRepository はUserRepositoryのポインタを生成する関数です
func NewTagRepository(sqlDB *sql.DB) *TagRepository {
	return &TagRepository{
		sqlDB: sqlDB,
	}
}

func (u TagRepository) Find(id string) (*model.Tag, error) {
	stmt, err := u.sqlDB.Prepare("SELECT * FROM tags WHERE id=?;")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tag model.Tag
	for rows.Next() {
		if err := rows.Scan(&tag.TagID, &tag.TagName); err != nil {
			return nil, err
		}
	}

	return &tag, nil
}

func (u TagRepository) FindAll() ([]*model.Tag, error) {
	rows, err := u.sqlDB.Query("SELECT * FROM tags")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res := make([]*model.Tag, 0)
	for rows.Next() {
		var tag model.Tag
		if err := rows.Scan(&tag.TagID, &tag.TagName); err != nil {
			return nil, err
		}
		res = append(res, &tag)
	}
	return res, nil
}

func (u TagRepository) Save(name string) (*model.Tag, error) {
	stmt, err := u.sqlDB.Prepare("INSERT INTO tags(tag_id, tag_name, created_at, updated_at) values(?, ?, ?, ?);")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var tag model.Tag
	tag.TagID = uuid.New().String()
	tag.TagName = name
	tag.UpdatedAt = time.Now()
	tag.CreatedAt = time.Now()

	_, err = stmt.Exec(tag.TagID, tag.TagName, tag.CreatedAt, tag.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &tag, err
}

func (u TagRepository) Delete(id string) error {
	stmt, err := u.sqlDB.Prepare("DELETE FROM tags WHERE id = ?;")
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(id); err != nil {
		return err
	}

	return nil
}
