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

	return echoContext.JSON(http.StatusOK, newResponseArray(*games))
}

func GetGameController(echoContext echo.Context) error {
	id := echoContext.Param("id")
	game, err := database.GetGameByID(id)
	if err != nil{
		return echoContext.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "not found",
		})
	}
	return echoContext.JSON(http.StatusOK, newResponse(*game))
}

func AddGameController(echoContext echo.Context) error {
	var gameReq GameRequest
	echoContext.Bind(&gameReq)
	result, err := database.PostGame(gameReq.toModel())
	if err != nil{
		return echoContext.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "bad request",
		})
	}

	return echoContext.JSON(http.StatusOK, newResponse(*result))
}

func UpdateGameController(echoContext echo.Context) error{
	var gameReq GameRequest
	echoContext.Bind(&gameReq)

	id := echoContext.Param("id")

	result, err := database.UpdateGame(id,gameReq.toModel())
	if err != nil{
		return echoContext.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "not found",
		})
	}

	return echoContext.JSON(http.StatusOK, newResponse(*result))
}

func DeleteGameController(echoContext echo.Context) error{
	id := echoContext.Param("id")
	game, err1 := database.GetGameByID(id)
	result, err2 := database.DeleteGame(id,*game)

	if err1 != nil{
		return echoContext.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "not found",
		})
	} else if err2 != nil{
		return echoContext.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "bad request",
		})
	}
	
	return echoContext.JSON(http.StatusOK,result)
}