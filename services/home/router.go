package home

import (
	"usermanagement/templates/home"

	"github.com/labstack/echo/v4"
)

func HomeRouter(e *echo.Echo) {
	e.GET("/", Home)
}

func Home(c echo.Context) error {
	user, err := c.Cookie("username")
	if err != nil {
		component := home.UI("")
		return component.Render(c.Request().Context(), c.Response())
	}
	component := home.UI(user.Value)
	return component.Render(c.Request().Context(), c.Response())
}
