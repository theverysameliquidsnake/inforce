package main

import (
	"context"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/theverysameliquidsnake/inforce/internal/database"
	"github.com/theverysameliquidsnake/inforce/internal/handler"
	"github.com/theverysameliquidsnake/inforce/internal/job"
	"github.com/theverysameliquidsnake/inforce/internal/repository"
	"github.com/theverysameliquidsnake/inforce/internal/service"
	ginprometheus "github.com/zsais/go-gin-prometheus"
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
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{os.Getenv("REACT_URL")},
		AllowMethods: []string{"GET", "POST"},
	}))

	prometheus := ginprometheus.NewWithConfig(ginprometheus.Config{
		Subsystem: "gin",
	})
	prometheus.Use(router)

	eventHandler.RegisterEventRoutes(router)

	if err := router.Run("0.0.0.0:8000"); err != nil {
		log.Fatalf("failed to start router: %v", err)
	}
}
