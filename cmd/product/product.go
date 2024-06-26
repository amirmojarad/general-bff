package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// Handler functions for orders
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

// Handler functions for products
func getProducts(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "GET /v1/products - List of products",
	})
}

func getProductById(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusOK, map[string]string{
		"message": "GET /v1/product/" + id + " - Product details",
	})
}

func createProduct(c echo.Context) error {
	return c.JSON(http.StatusCreated, map[string]string{
		"message": "POST /v1/product - Product created",
	})
}

func main() {
	// Create a new Echo instance
	e := echo.New()

	// Define routes for orders
	e.GET("/v1/orders", getOrders)
	e.GET("/v1/order/:id", getOrderById)
	e.POST("/v1/order", createOrder)

	// Define routes for products
	e.GET("/v1/products", getProducts)
	e.GET("/v1/product/:id", getProductById)
	e.POST("/v1/product", createProduct)

	// Start the server
	e.Logger.Fatal(e.Start(":9001"))
}
