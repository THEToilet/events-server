package gateway

import (
	"github.com/THEToilet/events-server/pkg/domain/repository"
	"github.com/gomodule/redigo/redis"
)

var _ repository.SessionRepository = &SessionRepository{}

type SessionRepository struct {
	redisConn redis.Conn
}

func NewSessionRepository(redisConn redis.Conn) *SessionRepository {
	return &SessionRepository{
		redisConn: redisConn,
	}
}

func (s SessionRepository) Find(sessionId string) (string, error) {
	rep, err := redis.String(s.redisConn.Do("GET", sessionId))
	return rep, err
}

func (s SessionRepository) Save(sessionId string, userId string) error {
	_, err := s.redisConn.Do("SET", sessionId, userId)
	return err
}
