package controllers

import (
	"discount-api/lib/database"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAllGameController(echoContext echo.Context) error {
	games, err := database.GetAllGame()
	if err != nil {
		return echoContext.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "bad request",
		})
	}

	return echoContext.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
		"games": games,
	})
}