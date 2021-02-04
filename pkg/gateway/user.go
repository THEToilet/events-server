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
	panic("implement me")
}

func (u UserRepository) FindAll() (*[]model.User, error) {
	panic("implement me")
}

func (u UserRepository) Save(id uuid.UUID, mail string) error {
	panic("implement me")
}

func (u UserRepository) Delete(id uuid.UUID) error {
	panic("implement me")
}

