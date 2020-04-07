package main

import (
	"Vegan_delivery_API/internal/hello/delivery"
	"Vegan_delivery_API/internal/hello/usecase"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"os"
)

func main() {
	// Echo instance
	e := echo.New()
	g := e.Group("/api/v1")

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// services
	helloService := usecase.NewService()

	// bind routes
	delivery.NewHandler(g, helloService)

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	port = ":" + port

	fmt.Printf("Starting server on port%s\n", port)
	log.Fatal(e.Start(port))
}
