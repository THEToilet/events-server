package gateway

import (
	"github.com/google/uuid"

	"../domain/model"
	"../domain/repository"
)

var _ repository.UserRepository = &UserRepository{}

// UserRepository は repository.UserRepository を満たす構造体です
type UserRepository struct {
	dbMap *gorp.DbMap
}

// NewUserRepository はUserRepositoryのポインタを生成する関数です
func NewUserRepository(dbMap *gorp.DbMap) *UserRepository {
	dbMap.AddTableWithName(userDTO{}, "users").SetKeys(false, "ID")
	return &UserRepository{dbMap: dbMap}
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

type userDTO struct {
	ID   string `db:"id"`
	Mail string `db:"mail"`
}
