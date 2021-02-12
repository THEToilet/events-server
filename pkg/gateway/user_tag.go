package gateway

import (
	"database/sql"
	"github.com/THEToilet/events-server/pkg/domain/model"
	"github.com/THEToilet/events-server/pkg/domain/repository"
	"github.com/google/uuid"
	"time"
)

var _ repository.UserTagRepository = &UserTagRepository{}

type UserTagRepository struct {
	sqlDB *sql.DB
}

func NewUserTagRepository(sqlDB *sql.DB) *UserTagRepository {
	return &UserTagRepository{
		sqlDB: sqlDB,
	}
}

func (u UserTagRepository) Find(id string) (*model.UserTag, error) {
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

	var tag model.UserTag
	for rows.Next() {
		if err := rows.Scan(&tag.TagID); err != nil {
			return nil, err
		}
	}

	return &tag, nil
}

func (u UserTagRepository) FindAll() ([]*model.UserTag, error) {
	rows, err := u.sqlDB.Query("SELECT * FROM tags")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res := make([]*model.UserTag, 0)
	for rows.Next() {
		var tag model.UserTag
		if err := rows.Scan(&tag.TagID); err != nil {
			return nil, err
		}
		res = append(res, &tag)
	}
	return res, nil
}

func (u UserTagRepository) Save(name string) (*model.UserTag, error) {
	stmt, err := u.sqlDB.Prepare("INSERT INTO tags(tag_id, tag_name, created_at, updated_at) values(?, ?, ?, ?);")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var tag model.UserTag
	tag.TagID = uuid.New().String()
	tag.UpdatedAt = time.Now()
	tag.CreatedAt = time.Now()

	_, err = stmt.Exec(tag.TagID, tag.CreatedAt, tag.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &tag, err
}

func (u UserTagRepository) Delete(id string) error {
	stmt, err := u.sqlDB.Prepare("DELETE FROM tags WHERE id = ?;")
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(id); err != nil {
		return err
	}

	return err
}
