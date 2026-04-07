package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type SummaryRepository struct {
	postgres *pgxpool.Pool
}

func NewSummaryRepository(postgres *pgxpool.Pool) *SummaryRepository {
	return &SummaryRepository{
		postgres: postgres,
	}
}

func (summaryRepo *SummaryRepository) Aggregate(ctx context.Context) error {
	query := `
		INSERT INTO user_event_summaries (user_id, events, start_time)
		SELECT user_id, COUNT(*), NOW() - INTERVAL '4 hours'
		FROM user_events
		WHERE timestamp >= NOW() - INTERVAL '4 hours'
		GROUP BY user_id
	`

	if _, err := summaryRepo.postgres.Exec(ctx, query); err != nil {
		return fmt.Errorf("aggregate: %w", err)
	}

	return nil
}
