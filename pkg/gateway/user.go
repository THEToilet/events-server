package gateway

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/THEToilet/events-server/pkg/domain/model"
	"github.com/THEToilet/events-server/pkg/domain/repository"
	"github.com/go-sql-driver/mysql"
)

var _ repository.UserRepository = &UserRepository{}

// UserRepository は repository.UserRepository を満たす構造体です
type UserRepository struct {
	sqlDB *sql.DB
}

// NewUserRepository はUserRepositoryのポインタを生成する関数です
func NewUserRepository(sqlDB *sql.DB) *UserRepository {
	return &UserRepository{
		sqlDB: sqlDB,
	}
}

func (u UserRepository) Find(id string) (*model.User, error) {
	stmt, err := u.sqlDB.Prepare("SELECT * FROM users WHERE user_id=?;")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	result := stmt.QueryRow(id)
	var user model.User
	if err := result.Scan(&user.UserID, &user.UserMail, &user.UserPassword); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("select user: %w", model.ErrUserNotFound)
		}
		return &user, fmt.Errorf("select user: %w", err)
	}

	return &user, nil
}

func (u UserRepository) FindAll() ([]*model.User, error) {
	rows, err := u.sqlDB.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res := make([]*model.User, 0)
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.UserID, &user.UserMail, &user.UserPassword); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, fmt.Errorf("select user: %w", model.ErrUserNotFound)
			}
			return res, fmt.Errorf("select user: %w", err)
		}
		res = append(res, &user)
	}
	return res, nil
}

func (u UserRepository) Save(user model.User) error {
	stmt, err := u.sqlDB.Prepare("INSERT INTO users(user_id, user_mail, user_password) values(?, ?, ?);")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.UserID, user.UserMail, user.UserPassword)
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == 1062 {
			return fmt.Errorf("save user: %w", model.ErrUserAlreadyExisted)
		}
		return fmt.Errorf("save user: %w", err)
	}
	return err
}

func (u UserRepository) Delete(id string) error {
	stmt, err := u.sqlDB.Prepare("DELETE FROM users WHERE user_id = ?;")
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(id); err != nil {
		return fmt.Errorf("delete: %w", err)
	}
	return err
}
