package usecase

import (
	"context"
	"github.com/THEToilet/events-server/pkg/domain/model"
	"github.com/THEToilet/events-server/pkg/domain/repository"
	"github.com/THEToilet/events-server/pkg/log"
	"go.uber.org/zap"
)

type TagUseCase struct {
	tagRepository repository.TagRepository
}

func NewTagUseCase(tagRepository repository.TagRepository) *TagUseCase {
	return &TagUseCase{
		tagRepository: tagRepository,
	}
}

func (t *TagUseCase) PostTag(ctx context.Context, name string) (*model.Tag, error) {
	logger := log.New()
	tag, err := t.tagRepository.Save(name)
	if err != nil {
		logger.Error("tag post failed", zap.Error(err))
		return nil, err
	}
	return tag, err
}

func (t *TagUseCase) GetTag(ctx context.Context) ([]*model.Tag, error) {
	logger := log.New()
	tags, err := t.tagRepository.FindAll()
	if err != nil {
		logger.Error("tag get failed", zap.Error(err))
		return nil, err
	}
	return tags, err
}
