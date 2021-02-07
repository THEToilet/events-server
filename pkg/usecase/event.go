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
	events, err := t.eventRepository.FindAll()
	return events, err
}

//PostEvent 新しくイベントを追加します
func (t *EventUseCase) PostEvent(ctx context.Context, id string) (*model.Event, error) {
	event, err := t.eventRepository.Save(id)
	return event, err
}

//PutEvent idで指定されたイベントを更新します。
func (t *EventUseCase) PutEvent(ctx context.Context, id string) (*model.Event, error) {
	event, err := t.eventRepository.Update(id)
	return event, err
}

//DeleteEvent idで指定されたイベントを削除します。
func (t *EventUseCase) DeleteEvent(ctx context.Context, id string) error {
	err := t.eventRepository.Delete(id)
	return err
}
