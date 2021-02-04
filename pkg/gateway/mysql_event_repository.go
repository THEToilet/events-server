package gateway

import (
	"database/sql"
)

type Event struct {
	Db         sql.DB
	EventID    string `json:"event_id"`
	PostedUser string `json:"posted_user"`
	PostedDate string `json:"posted_date"`
	Message    string `json:"message"`
}

func (e *Event) find(event_id string) (err error) {
	err = e.Db.QueryRow("select id, content, author from posts where id = $1", event_id).Scan(&e.EventID, &e.Message, &e.PostedDate)
	return
}

func (e *Event) findAll() (err error) {
	return
}

func (e *Event) save() (err error) {
	statement := "insert into events (content, author) values ($1, $2) returning id"
	stmt, err := e.Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(post.Content, post.Author).Scan(&e.EventID)
	return
}

func (e *Event) update() (err error) {
	_, err = e.Db.Exec("update events set posted_user = $2, posted_date = $3, message = $4 where event_id = $1", e.EventID, e.PostedDate, e.PostedDate, e.Message)
	return
}

func (e *Event) delete() (err error) {
	_, err = e.Db.Exec("delete from posts where id = $1", e.EventID)
	return
}
