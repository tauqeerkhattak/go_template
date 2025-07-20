package controllers

import (
	"database/sql"
	"net/http"
	"template/models"
	"template/utils"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	DB *sql.DB
}

func (ac *AuthController) Register(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	hashedPassword, _ := utils.HashPassword(user.Password)
	_, err := ac.DB.Exec("INSERT INTO users (name, email, password) VALUES (?, ?, ?)", user.Name, user.Email, hashedPassword)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "User already exists or DB error"})
	}

	return c.JSON(http.StatusCreated, echo.Map{"message": "User registered successfully"})
}

func (ac *AuthController) Login(c echo.Context) error {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	var user models.User
	row := ac.DB.QueryRow("SELECT id, password FROM users WHERE email = ?", input.Email)
	err := row.Scan(&user.ID, &user.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid email or password"})
	}

	if !utils.CheckPasswordHash(input.Password, user.Password) {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid email or password"})
	}

	token, _ := utils.GenerateJWT(user.ID)
	return c.JSON(http.StatusOK, echo.Map{"token": token})
}
