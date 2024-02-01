package auth

import (
	"net/http"
	"time"
	"usermanagement/templates/auth"

	"github.com/labstack/echo/v4"
)

const ADMIN = "Admin"
const USER = "User"

type LoginFormData struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type SignupFormData struct {
	Username  string `json:"username" form:"username"`
	Email     string `json:"email" form:"email"`
	Password  string `json:"password" form:"password"`
	Workspace string `json:"workspace" form:"workspace"`
}

type User struct {
	Username  string `json:"username" form:"username"`
	Email     string `json:"email" form:"email"`
	Password  string `json:"_"`
	Workspace string `json:"workspace" form:"workspace"`
	Role      string `json:"role" form:"role"`
}

var Users []User = []User{
	{
		Username:  "firstuser",
		Email:     "email@email.com",
		Password:  "123456",
		Workspace: "myworkspace",
		Role:      ADMIN,
	},
}

func AuthRouter(e *echo.Echo) {
	e.GET("/login", Login)
	e.GET("/logout", LogoutHandler)
	e.POST("/login", LoginHandler)
	e.GET("/signup", Signup)
}

func Login(c echo.Context) error {
	component := auth.LoginUI()
	return component.Render(c.Request().Context(), c.Response())
}

func LoginHandler(c echo.Context) error {
	var loginData LoginFormData

	err := c.Bind(&loginData)
	if err != nil {
		return c.HTML(http.StatusBadRequest, "form data handling error")
	}

	for _, user := range Users {
		if user.Email == loginData.Email && user.Password == loginData.Password {
			// Set a cookie
			cookie := new(http.Cookie)
			cookie.Name = "username"
			cookie.Value = user.Username
			cookie.Expires = time.Now().Add(24 * time.Hour)
			c.SetCookie(cookie)
			return c.Redirect(http.StatusSeeOther, "/")
		}
	}

	return c.HTML(http.StatusBadRequest, "wrong email or password")
}

func LogoutHandler(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "username"
	cookie.Expires = time.Now().AddDate(2000, 0, 0)
	c.SetCookie(cookie)
	return c.Redirect(http.StatusSeeOther, "/")
}

func Signup(c echo.Context) error {
	component := auth.SignupUI()
	return component.Render(c.Request().Context(), c.Response())
}

func SignupHandler(c echo.Context) error {
	var loginData LoginFormData

	err := c.Bind(&loginData)
	if err != nil {
		return c.HTML(http.StatusBadRequest, "form data handling error")
	}

	return nil
}
