package pkg

import (
	"github.com/labstack/echo"
	"net/http"
)

func Routes(e *echo.Echo) {
	// Health Page
	e.GET("/health", health)

	//v1User := e.Group("/api/v1/users")
	//model.UserRouter(v1User)
}

func health(c echo.Context) error {
	return c.String(http.StatusOK, "I am live!")
}
