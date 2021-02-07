package usecase

import (
	"context"
	"github.com/THEToilet/events-server/pkg/domain/model"
	"github.com/THEToilet/events-server/pkg/domain/repository"
)

type TagUseCase struct {
	tagRepository repository.TagRepository
}

func NewTagUseCase(tagRepository repository.TagRepository) *TagUseCase {
	return &TagUseCase{
		tagRepository: tagRepository,
	}
}

func (t *TagUseCase) PostTag(ctx context.Context) error {

}

func (t *TagUseCase) GetTagList(ctx context.Context) ([]*model.Tag, error) {

}
