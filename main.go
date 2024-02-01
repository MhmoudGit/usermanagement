package main

import (
	"net/http"
	"time"
	"usermanagement/templates/auth"
	"usermanagement/templates/home"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.Recover())

	// Serving static files (css, html, images)
	e.Static("/static", "static")

	e.GET("/", Home)
	e.GET("/login", Login)
	e.GET("/logout", LogoutHandler)
	e.POST("/login", LoginHandler)

	e.Start("localhost:8000")
}

type LoginFormData struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var Users []LoginFormData = []LoginFormData{
	{Email: "email@email.com", Password: "123456"},
}

func Home(c echo.Context) error {
	user, err := c.Cookie("userEmail")
	if err != nil {
		component := home.UI("")
		return component.Render(c.Request().Context(), c.Response())
	}
	component := home.UI(user.Value)
	return component.Render(c.Request().Context(), c.Response())
}

func Login(c echo.Context) error {
	component := auth.LoginUI()
	return component.Render(c.Request().Context(), c.Response())
}

func LoginHandler(c echo.Context) error {
	var loginData LoginFormData

	err := c.Bind(&loginData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "form data handling error")
	}

	for _, user := range Users {
		if user.Email == loginData.Email && user.Password == loginData.Password {
			// Set a cookie
			cookie := new(http.Cookie)
			cookie.Name = "userEmail"
			cookie.Value = user.Email
			cookie.Expires = time.Now().Add(24 * time.Hour)
			c.SetCookie(cookie)
			return c.Redirect(http.StatusSeeOther, "/")
		}
	}

	return c.JSON(http.StatusBadRequest, "wrong email or password")
}

func LogoutHandler(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "userEmail"
	cookie.Expires = time.Now().AddDate(2000, 0, 0)
	c.SetCookie(cookie)
	return c.Redirect(http.StatusSeeOther, "/")
}
