package gateway

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/THEToilet/events-server/pkg/domain/model"
	"github.com/THEToilet/events-server/pkg/domain/repository"
	"github.com/go-sql-driver/mysql"
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

func (u UserTagRepository) FindAll(eventID string) ([]*model.UserTag, error) {
	rows, err := u.sqlDB.Query("SELECT * FROM users_tags WHERE event_id = ?", eventID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res := make([]*model.UserTag, 0)
	for rows.Next() {
		var tag model.UserTag
		if err := rows.Scan(&tag.EventID, &tag.TagID, &tag.CreatedAt, &tag.UpdatedAt); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, fmt.Errorf("select user_tags: %w", model.ErrUserTagNotFound)
			}
			return res, fmt.Errorf("select user_tags: %w", err)
		}
		res = append(res, &tag)
	}
	return res, nil
}

func (u UserTagRepository) Save(tag *model.UserTag) error {
	stmt, err := u.sqlDB.Prepare("INSERT INTO users_tags(event_id, tag_id, created_at, updated_at) values(?, ?, ?, ?);")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(tag.EventID, tag.TagID, tag.CreatedAt, tag.UpdatedAt)
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == 1062 {
			return fmt.Errorf("save user_tags: %w", model.ErrTagAlreadyExisted)
		}
		return fmt.Errorf("save user_tags: %w", err)
	}
	return err
}

func (u UserTagRepository) Delete(eventID string, tagID string) error {
	stmt, err := u.sqlDB.Prepare("DELETE FROM users_tags WHERE event_id = ? && tag_id = ?;")
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(eventID, tagID); err != nil {
		return fmt.Errorf("delete: %w", err)
	}

	return err
}
