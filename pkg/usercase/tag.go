package usercase

import (
	"../domain/model"
	"../domain/repository"
	"fmt"
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
	tag, err := t.tagRepository.Find(uuid.New())
	if err != nil {
		fmt.Errorf("unko")
	}
	return tag, err
}
