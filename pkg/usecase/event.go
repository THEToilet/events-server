package usecase

import (
	"context"
	"github.com/THEToilet/events-server/pkg/domain/model"
	"github.com/THEToilet/events-server/pkg/domain/repository"
)

type EventUseCase struct {
	eventRepository repository.EventRepository
}

func NewEventUseCase(eventRepository repository.EventRepository) *EventUseCase {
	return &EventUseCase{
		eventRepository: eventRepository,
	}
}

//GetEvent イベントの一覧を取得します
func (t *EventUseCase) GetEvent(ctx context.Context) ([]*model.Event, error) {

}

//PostEvent 新しくイベントを追加します
func (t *EventUseCase) PostEvent(ctx context.Context) (*model.Event, error) {
}

//PutEvent idで指定されたイベントを更新します。
func (t *EventUseCase) PutEvent(ctx context.Context, id string) (*model.Event, error) {
}

//DeleteEvent idで指定されたイベントを削除します。
func (t *EventUseCase) DeleteEvent(ctx context.Context, id string) (*model.Event, error) {
}
