package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"template/config"
	"template/database"
	"template/routes"
)

func main() {
	config.LoadEnv()
	db := database.InitDB()
	defer db.Close()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	routes.SetupRoutes(e, db)

	port := config.Env("PORT", "8080")
	log.Fatal(e.Start(":" + port))
}
