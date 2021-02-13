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
	stmt, err := u.sqlDB.Prepare("SELECT * FROM events where events_id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var event model.Event
	for rows.Next() {
		if err := rows.Scan(&event.EventID, &event.DeadLine, &event.PostedUser, &event.Description, &event.CreatedAt, &event.UpdatedAt); err != nil {
			return nil, err
		}
	}
	//event.Tag = tagsRes
	return &event, nil
}

func (u EventRepository) FindAll() ([]*model.Event, error) {
	rows, err := u.sqlDB.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res := make([]*model.Event, 0)
	for rows.Next() {
		var event model.Event
		if err := rows.Scan(&event.EventID); err != nil {
			return nil, err
		}
		res = append(res, &event)
	}
	return res, nil
}

func (u EventRepository) Save(id string) (*model.Event, error) {
	stmt, err := u.sqlDB.Prepare("INSERT INTO users(id, mail) values(?, ?);")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	insertResult, err := stmt.Exec(id)
	if err != nil {
		return nil, err
	}

	userID, err := insertResult.LastInsertId()
	if err != nil {
		return nil, err
	}
	fmt.Println(userID)

	return nil, nil
}

func (u EventRepository) Delete(id string) error {
	stmt, err := u.sqlDB.Prepare("DELETE FROM users WHERE id = ?;")
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(id); err != nil {
		return err
	}

	return nil
}

func (u EventRepository) Update(id string) (*model.Event, error) {
	stmt, err := u.sqlDB.Prepare("INSERT INTO users(id, mail) values(?, ?);")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	insertResult, err := stmt.Exec(id)
	if err != nil {
		return nil, err
	}

	userID, err := insertResult.LastInsertId()
	if err != nil {
		return nil, err
	}
	fmt.Println(userID)

	return nil, nil
}
