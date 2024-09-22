package handler

import (
	"net/http"
	"registro/kafka"

	"github.com/labstack/echo/v4"
)

type EventRequest struct {
	Topic   string `json:"topic"`
	Message string `json:"message"`
}

func ProduceEvent(c echo.Context) error {
	var event EventRequest
	if err := c.Bind(&event); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	err := kafka.ProduceEvent(event.Topic, []byte(event.Message))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to produce event"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Event produced successfully"})
}

func GetEvents(c echo.Context) error {
	// Placeholder for fetching events from the blockchain
	return c.JSON(http.StatusOK, map[string]string{"message": "List of events..."})
}
