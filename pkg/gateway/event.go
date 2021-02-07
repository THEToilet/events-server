package gateway

import (
	"database/sql"
	"fmt"
	"github.com/THEToilet/events-server/pkg/domain/model"
	"github.com/THEToilet/events-server/pkg/domain/repository"
)

var _ repository.EventRepository = &EventRepository{}

// UserRepository は repository.UserRepository を満たす構造体です
type EventRepository struct {
	sqlDB *sql.DB
}

// NewUserRepository はUserRepositoryのポインタを生成する関数です
func NewEventRepository(sqlDB *sql.DB) *EventRepository {
	return &EventRepository{
		sqlDB: sqlDB,
	}
}

func (u EventRepository) Find(id string) (*model.Event, error) {
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

func (u EventRepository) FindAll() (*[]*model.Event, error) {
	rows, err := u.sqlDB.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res := make([]*model.User, 0)
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Mail); err != nil {
			return nil, err
		}
		res = append(res, &user)
	}
	return &res, nil
}

func (u EventRepository) Save(id string) error {
	stmt, err := u.sqlDB.Prepare("INSERT INTO users(id, mail) values(?, ?);")
	if err != nil {
		return err
	}
	defer stmt.Close()

	insertResult, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	userID, err := insertResult.LastInsertId()
	if err != nil {
		return err
	}
	fmt.Println(userID)

	return nil
}

func (u EventRepository) Delete(id string) error {
	stmt, err := u.sqlDB.Prepare("DELETE FROM users WHERE id + ?;")
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(id); err != nil {
		return err
	}

	return nil
}

func (u EventRepository) Update(id string) error {
	stmt, err := u.sqlDB.Prepare("INSERT INTO users(id, mail) values(?, ?);")
	if err != nil {
		return err
	}
	defer stmt.Close()

	insertResult, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	userID, err := insertResult.LastInsertId()
	if err != nil {
		return err
	}
	fmt.Println(userID)

	return nil
}
