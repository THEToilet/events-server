package usercase

import (
	"../domain/model"
	"../domain/repository"
)

type EventUseCase struct {
	eventRepository repository.EventRepository
}

func NewEventUseCase(eventRepository repository.EventRepository) *EventUseCase {
	return &EventUseCase{
		eventRepository: eventRepository,
	}
}

func (t *EventUseCase) GetEvent() (*model.Event, error) {
}
