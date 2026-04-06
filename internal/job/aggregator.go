package job

import (
	"context"
	"log"
	"time"

	"github.com/robfig/cron/v3"
	"github.com/theverysameliquidsnake/inforce/internal/service"
)

func StartAggregationBackgroundJob(ctx context.Context, summaryService *service.SummaryService) {
	cron := cron.New(cron.WithLocation(time.UTC))

	_, err := cron.AddFunc("@every 4h", func() {
		if err := summaryService.Aggregate(ctx); err != nil {
			log.Printf("background job failed: %v", err)
		}
	})
	if err != nil {
		log.Fatalf("failed to schedule background job: %v", err)
	}

	cron.Start()
}
