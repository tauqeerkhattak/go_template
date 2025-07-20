package routes

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"template/controllers"
	"template/middlewares"
)

func SetupRoutes(e *echo.Echo, db *sql.DB) {
	auth := controllers.AuthController{DB: db}

	e.POST("/api/register", auth.Register)
	e.POST("/api/login", auth.Login)

	protected := e.Group("/api")
	protected.Use(middlewares.JWTMiddleware)
	protected.GET("/profile", func(c echo.Context) error {
		return c.JSON(200, echo.Map{"message": "This is a protected route"})
	})
}
