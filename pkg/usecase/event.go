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

func (t *EventUseCase) GetEvent(context.Context) (*model.Event, error) {
}
func (t *EventUseCase) PostEvent(context.Context) (*model.Event, error) {
}
func (t *EventUseCase) PutEvent(context.Context) (*model.Event, error) {
}
func (t *EventUseCase) DeleteEvent(context.Context) (*model.Event, error) {
}
