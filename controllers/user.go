package controllers

import (
	"discount-api/lib/database"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAllUserController(echoContext echo.Context) error {
	users, err := database.GetAllUser()
	if err != nil {
		return echoContext.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "bad request",
		})
	}

	return echoContext.JSON(http.StatusOK, newResponseArray2(*users))
}

func GetUserController(echoContext echo.Context) error {
	id := echoContext.Param("id")
	user, err := database.GetUserByID(id)
	if err != nil{
		return echoContext.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "not found",
		})
	}
	return echoContext.JSON(http.StatusOK, newResponse2(*user))
}

func AddUserController(echoContext echo.Context) error {
	var userReq UserRequest
	echoContext.Bind(&userReq)
	result, err := database.PostUser(userReq.toModel())
	if err != nil{
		return echoContext.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "bad request",
		})
	}

	return echoContext.JSON(http.StatusOK, newResponse2(*result))
}

func UpdateUserController(echoContext echo.Context) error{
	var userReq UserRequest
	echoContext.Bind(&userReq)

	id := echoContext.Param("id")

	result, err := database.UpdateUser(id,userReq.toModel())
	if err != nil{
		return echoContext.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "not found",
		})
	}

	return echoContext.JSON(http.StatusOK, newResponse2(*result))
}

func DeleteUserController(echoContext echo.Context) error{
	id := echoContext.Param("id")
	user, err1 := database.GetUserByID(id)
	result, err2 := database.DeleteUser(id,*user)

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