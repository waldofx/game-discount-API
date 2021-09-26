package main

import (
	"discount-api/config"
	ownMid "discount-api/middleware"
	"discount-api/routes"
)

func main() {
	config.InitDB()
	e := routes.NewRoutes()
	ownMid.LogMiddlewareInit(e)
	e.Start("localhost:8000")
}