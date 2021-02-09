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
	stmt, err := u.sqlDB.Prepare("SELECT users_tags.tag_id, tags.tag_id, users_tags.created_at, users_tags.updated_at FROM tags JOIN users_tags ON users_tags.tag_id = tags.tag_id WHERE users_tags.events_id = ?;")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tagsRes := make([]*model.Tag, 0)
	for rows.Next() {
		var tag model.Tag
		if err := rows.Scan(&tag.ID, &tag.Name, &tag.CreatedAt, &tag.UpdatedAt); err != nil {
			return nil, err
		}
		tagsRes = append(tagsRes, &tag)
	}

	stmt1, err := u.sqlDB.Prepare("SELECT * FROM events where events_id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt1.Close()

	rows1, err := stmt1.Query(id)
	if err != nil {
		return nil, err
	}
	defer rows1.Close()

	var event model.Event
	for rows1.Next() {
		if err := rows1.Scan(&event.ID, &event.DeadLine, &event.PostedUser, &event.Description, &event.CreatedAt, &event.UpdatedAt); err != nil {
			return nil, err
		}
	}
	event.Tag = tagsRes
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
		var user model.Event
		if err := rows.Scan(&user.ID); err != nil {
			return nil, err
		}
		res = append(res, &user)
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
