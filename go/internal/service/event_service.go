package service

import (
	"context"
	"fmt"

	"github.com/theverysameliquidsnake/inforce/internal/dto"
	"github.com/theverysameliquidsnake/inforce/internal/model"
	"github.com/theverysameliquidsnake/inforce/internal/repository"
)

type EventService struct {
	eventRepo *repository.EventRepository
}

func NewEventService(eventRepo *repository.EventRepository) *EventService {
	return &EventService{
		eventRepo: eventRepo,
	}
}

func (eventService *EventService) Create(ctx context.Context, event *dto.CreateEvent) (int, error) {
	id, err := eventService.eventRepo.Create(ctx, event)
	if err != nil {
		return 0, fmt.Errorf("create event service: %w", err)
	}

	return id, nil
}

func (eventService *EventService) Find(ctx context.Context, filter *dto.FilterEvent) ([]*model.Event, error) {
	events, err := eventService.eventRepo.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("find event service: %w", err)
	}

	return events, nil
}
