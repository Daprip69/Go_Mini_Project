package main

import (
	"Go_Mini_Project/config"
	"Go_Mini_Project/middleware"
	"Go_Mini_Project/route"

	"github.com/labstack/echo/v4"
)

func main() {
	db := config.InitDB()
	e := echo.New()
	middleware.Logmiddleware(e)

	route.NewRoute(e, db)

	e.Logger.Fatal(e.Start(":8080"))
}
