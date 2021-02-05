package usecase

import (
	"fmt"
	"github.com/THEToilet/events-server/pkg/domain/model"
	"github.com/THEToilet/events-server/pkg/domain/repository"
	"github.com/google/uuid"
)

type TagUseCase struct {
	tagRepository repository.TagRepository
}

func NewTagUseCase(tagRepository repository.TagRepository) *TagUseCase{
	return &TagUseCase{
		tagRepository: tagRepository,
	}
}

func (t *TagUseCase) GetTag() (*model.Tag, error) {
	tag, err := t.tagRepository.Find(uuid.New().String())
	if err != nil {
		fmt.Errorf("unko")
	}
	return tag, err
}
