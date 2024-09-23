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
	return c.JSON(http.StatusOK, kafka.Ledger.GetBlockchain())
}

// ValidateBlockchain checks the integrity of the ledger.
func ValidateBlockchain(c echo.Context) error {
	isValid := kafka.Ledger.ValidateLedger()
	return c.JSON(http.StatusOK, map[string]bool{"isValid": isValid})
}
