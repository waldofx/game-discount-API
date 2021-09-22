package main

import (
	"discount-api/config"
	ownMid "discount-api/middleware"
	"discount-api/routes"
)

func main() {
	config.InitDB()
	config.InitMigrate()
	e := routes.NewRoutes()
	ownMid.LogMiddlewareInit(e)
	e.Start(":8000")
}