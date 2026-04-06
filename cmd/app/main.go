package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/theverysameliquidsnake/inforce/internal/database"
	"github.com/theverysameliquidsnake/inforce/internal/handler"
	"github.com/theverysameliquidsnake/inforce/internal/job"
	"github.com/theverysameliquidsnake/inforce/internal/repository"
	"github.com/theverysameliquidsnake/inforce/internal/service"
)

func init() {
	_ = godotenv.Load()
}

func main() {
	postgres, err := database.NewPostgresPool()
	if err != nil {
		log.Fatalf("failed to create postgres pool: %v", err)
	}
	defer postgres.Close()

	eventRepo := repository.NewEventRepository(postgres)
	eventService := service.NewEventService(eventRepo)
	eventHandler := handler.NewEventHandler(eventService)

	summaryRepo := repository.NewSummaryRepository(postgres)
	summaryService := service.NewSummaryService(summaryRepo)
	go job.StartAggregationBackgroundJob(context.Background(), summaryService)

	router := gin.Default()

	eventHandler.RegisterEventRoutes(router)

	if err := router.Run(":8000"); err != nil {
		log.Fatalf("failed to start router: %v", err)
	}
}
