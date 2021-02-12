package gateway

import (
	"errors"
	"fmt"
	"github.com/THEToilet/events-server/pkg/domain/model"
	"github.com/THEToilet/events-server/pkg/domain/repository"
	"github.com/gomodule/redigo/redis"
)

var _ repository.SessionRepository = &SessionRepository{}
var single = make(chan bool, 1)

type SessionRepository struct {
	redisConn redis.Conn
}

func NewSessionRepository(redisConn redis.Conn) *SessionRepository {
	return &SessionRepository{
		redisConn: redisConn,
	}
}

func (s SessionRepository) Find(sessionId string) (string, error) {
	single <- false
	rep, err := redis.String(s.redisConn.Do("GET", sessionId))
	<-single
	if err != nil {
		if errors.Is(err, redis.ErrNil) {
			return "", fmt.Errorf("get session: %w", model.ErrSessionNotFound)
		}
		return "", fmt.Errorf("get session: %w", err)
	}
	return rep, err
}

func (s SessionRepository) Save(sessionId string, userId string) error {
	single <- false
	_, err := s.redisConn.Do("SET", sessionId, userId)
	<-single
	if err != nil {
		return fmt.Errorf("get session: %w", err)
	}
	return err
}
