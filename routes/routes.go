package routes

import (
	"discount-api/constants"
	"discount-api/controllers"

	"github.com/labstack/echo/middleware"
	"github.com/labstack/echo/v4"
)

func NewRoutes() *echo.Echo{
	e := echo.New()
	//game
	e.GET("/game",controllers.GetAllGameController)
	e.GET("/game/:id",controllers.GetGameController)
	e.PUT("/game/:id",controllers.UpdateGameController)
	e.DELETE("/game/:id",controllers.DeleteGameController)
	e.POST("/game",controllers.AddGameController)

	//user
	//e.GET("/user",controllers.GetAllUserController)
	e.GET("/user/:id",controllers.GetUserController)
	e.PUT("/user/:id",controllers.UpdateUserController)
	e.DELETE("/user/:id",controllers.DeleteUserController)
	e.POST("/user",controllers.AddUserController)


	//jwt
	//e.Pre(middleware.RemoveTrailingSlash())
	jwt := middleware.JWT([]byte(constants.JWT_SECRET))
	e.GET("/user",controllers.GetAllUserController, jwt)
	return e
}