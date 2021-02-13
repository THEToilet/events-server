package gateway

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/THEToilet/events-server/pkg/domain/model"
	"github.com/THEToilet/events-server/pkg/domain/repository"
	"github.com/go-sql-driver/mysql"
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
	stmt, err := u.sqlDB.Prepare("SELECT * FROM tags WHERE tag_id=?;")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	result := stmt.QueryRow(id)
	var tag model.Tag
	if err := result.Scan(&tag.TagID, &tag.TagName, &tag.CreatedAt, &tag.UpdatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("select tags: %w", model.ErrTagNotFound)
		}
		return &tag, fmt.Errorf("select tags: %w", err)
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
		if err := rows.Scan(&tag.TagID, &tag.TagName, &tag.CreatedAt, &tag.UpdatedAt); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, fmt.Errorf("select tags: %w", model.ErrTagNotFound)
			}
			return res, fmt.Errorf("select tags: %w", err)
		}
		res = append(res, &tag)
	}
	return res, nil
}

func (u TagRepository) Save(tag *model.Tag) error {
	stmt, err := u.sqlDB.Prepare("INSERT INTO tags(tag_id, tag_name, created_at, updated_at) values(?, ?, ?, ?);")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(tag.TagID, tag.TagName, tag.CreatedAt, tag.UpdatedAt)
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == 1062 {
			return fmt.Errorf("save tags: %w", model.ErrTagAlreadyExisted)
		}
		return fmt.Errorf("save tags: %w", err)
	}
	return err
}

func (u TagRepository) Delete(id string) error {
	stmt, err := u.sqlDB.Prepare("DELETE FROM tags WHERE tag_id = ?;")
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(id); err != nil {
		return fmt.Errorf("delete: %w", err)
	}

	return nil
}
