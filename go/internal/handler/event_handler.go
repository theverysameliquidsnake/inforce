package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/theverysameliquidsnake/inforce/internal/dto"
	"github.com/theverysameliquidsnake/inforce/internal/service"
)

type EventHandler struct {
	eventService *service.EventService
}

func NewEventHandler(eventService *service.EventService) *EventHandler {
	return &EventHandler{
		eventService: eventService,
	}
}

func (eventHandler *EventHandler) RegisterEventRoutes(router *gin.Engine) {
	router.POST("/event", func(ctx *gin.Context) {
		var params dto.CreateEvent
		if err := ctx.ShouldBindJSON(&params); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		id, err := eventHandler.eventService.Create(ctx.Request.Context(), &params)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{
			"event_id": id,
		})
	})

	router.GET("/event", func(ctx *gin.Context) {
		var params dto.FilterEvent
		if err := ctx.ShouldBindQuery(&params); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		events, err := eventHandler.eventService.Find(ctx.Request.Context(), &params)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"events": events,
		})
	})
}
