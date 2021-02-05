package gateway

import (
	"database/sql"
	"github.com/google/uuid"

	"github.com/THEToilet/events-server/pkg/domain/model"
	"github.com/THEToilet/events-server/pkg/domain/repository"
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

func (u UserRepository) Find(id uuid.UUID) (*model.User, error) {
	stmt, err := u.sqlDB.Prepare("SELECT * FROM users WHERE id=?;")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var user model.User
	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Mail); err != nil {
			return nil, err
		}
	}

	return &user, nil
}

func (u UserRepository) FindAll() (*[]model.User, error) {
	stmt, err := u.sqlDB.Prepare(querySelectUserByStatus)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Date, &user.Status); err != nil {
			return nil, parseError(err)
		}
		res = append(res, user)
	}

	return res, nil
}

func (u UserRepository) Save(id uuid.UUID, mail string) error {
}

func (u UserRepository) Delete(id uuid.UUID) error {
}
