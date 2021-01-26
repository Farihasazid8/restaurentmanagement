package pkg

import (
	"github.com/labstack/echo"
	"net/http"
	v1 "restaurentManagement/pkg/v1/model"
)

func Routes(e *echo.Echo) {
	// Health Page
	e.GET("/health", health)

	v1Employee := e.Group("/api/v1/employee")
	v1.EmployeeRouter(v1Employee)

	v1Customer := e.Group("/api/v1/customer")
	v1.CustomerRouter(v1Customer)

	v1Ingredient := e.Group("/api/v1/ingredient")
	v1.IngredientRouter(v1Ingredient)

	v1Dishes := e.Group("/api/v1/dish")
	v1.DishRouter(v1Dishes)

	v1Bill := e.Group("/api/v1/bill")
	v1.BillRouter(v1Bill)
}

func health(c echo.Context) error {
	return c.String(http.StatusOK, "I am live!")
}
