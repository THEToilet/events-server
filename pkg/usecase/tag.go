package usecase

import (
	"context"
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


func (t *TagUseCase) PostTag(ctx context.Context) (interface{}, interface{}) {

}

func (t *TagUseCase) GetTagList(ctx context.Context) (interface{}, interface{}) {

}
