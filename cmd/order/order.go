package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// Handler functions
func getOrders(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "GET /v1/orders - List of orders",
	})
}

func getOrderById(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusOK, map[string]string{
		"message": "GET /v1/order/" + id + " - Order details",
	})
}

func createOrder(c echo.Context) error {
	return c.JSON(http.StatusCreated, map[string]string{
		"message": "POST /v1/order - Order created",
	})
}

func main() {
	// Create a new Echo instance
	e := echo.New()

	// Define routes
	e.GET("/v1/orders", getOrders)
	e.GET("/v1/order/:id", getOrderById)
	e.POST("/v1/order", createOrder)

	// Start the server
	e.Logger.Fatal(e.Start(":9000"))
}
