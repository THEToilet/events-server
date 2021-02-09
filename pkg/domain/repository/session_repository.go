package repository

type SessionRepository interface {
	Find(sessionId string) (string, error)
	Save(sessionId string, userId string) error
}
