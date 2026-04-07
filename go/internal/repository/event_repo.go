package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/theverysameliquidsnake/inforce/internal/dto"
	"github.com/theverysameliquidsnake/inforce/internal/model"
)

type EventRepository struct {
	postgres *pgxpool.Pool
}

func NewEventRepository(postgres *pgxpool.Pool) *EventRepository {
	return &EventRepository{
		postgres: postgres,
	}
}

func (eventRepo *EventRepository) Create(ctx context.Context, event *dto.CreateEvent) (int, error) {
	query := `
		INSERT INTO user_events (user_id, action, metadata)
		VALUES ($1, $2, $3)
		RETURNING id
	`

	var metadata []byte
	var err error
	if event.Metadata == nil {
		metadata = nil
	} else {
		metadata, err = json.Marshal(event.Metadata)
		if err != nil {
			return 0, fmt.Errorf("marshal metadata: %w", err)
		}
	}

	var id int
	err = eventRepo.postgres.QueryRow(ctx, query, event.UserId, event.Action, metadata).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("insert event: %w", err)
	}

	return id, nil
}

func (eventRepo *EventRepository) Find(ctx context.Context, filter *dto.FilterEvent) ([]*model.Event, error) {
	query := `
		SELECT id, user_id, action, timestamp, metadata
		FROM user_events
	`

	var where []string
	var args []any

	if filter.UserId != nil {
		where = append(where, fmt.Sprintf("user_id = $%d", len(where)+1))
		args = append(args, *filter.UserId)
	}

	if filter.StartTime != nil {
		where = append(where, fmt.Sprintf("timestamp >= $%d", len(where)+1))
		args = append(args, *filter.StartTime)
	}

	if filter.EndTime != nil {
		where = append(where, fmt.Sprintf("timestamp <= $%d", len(where)+1))
		args = append(args, *filter.EndTime)
	}

	if len(where) > 0 {
		query += " WHERE " + strings.Join(where, " AND ")
	}

	query += " ORDER BY timestamp DESC"

	rows, err := eventRepo.postgres.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("query events: %w", err)
	}
	defer rows.Close()

	var events []*model.Event
	for rows.Next() {
		var event model.Event
		var metadata []byte

		if err := rows.Scan(&event.Id, &event.UserId, &event.Action, &event.Timestamp, &metadata); err != nil {
			return nil, fmt.Errorf("scan event: %w", err)
		}

		if len(metadata) > 0 {
			if err := json.Unmarshal(metadata, &event.Metadata); err != nil {
				return nil, fmt.Errorf("unmarshal metadata: %w", err)
			}
		}

		events = append(events, &event)
	}

	return events, nil
}
