package main

import (
	"fmt"
	dishhttp "github.com/efimovad/Vegan_delivery_API/internal/app/dish/delivery/http"
	"github.com/efimovad/Vegan_delivery_API/internal/app/hello/delivery"
	"github.com/efimovad/Vegan_delivery_API/internal/app/hello/usecase"
	placehttp "github.com/efimovad/Vegan_delivery_API/internal/app/place/delivery/http"
	profilehttp "github.com/efimovad/Vegan_delivery_API/internal/app/profile/delivery/http"
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
	dishhttp.NewHandler(g)
	placehttp.NewHandler(g)
	profilehttp.NewHandler(g)


	port := os.Getenv("PORT")
	fmt.Printf("env var PORT: %s\n", port)
	if port == "" {
		port = "8080"
	}
	port = ":" + port

	fmt.Printf("Starting server on port%s\n", port)
	log.Fatal(e.Start(port))
}
