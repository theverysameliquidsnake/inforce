package service

import (
	"context"
	"fmt"

	"github.com/theverysameliquidsnake/inforce/internal/repository"
)

type SummaryService struct {
	summaryRepo *repository.SummaryRepository
}

func NewSummaryService(summaryRepo *repository.SummaryRepository) *SummaryService {
	return &SummaryService{
		summaryRepo: summaryRepo,
	}
}

func (summaryService *SummaryService) Aggregate(ctx context.Context) error {
	if err := summaryService.summaryRepo.Aggregate(ctx); err != nil {
		return fmt.Errorf("aggregate service: %w", err)
	}

	return nil
}
