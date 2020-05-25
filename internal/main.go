package main

import (
	"fmt"
	"github.com/efimovad/Vegan_delivery_API/internal/app"
	"log"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	url := os.Getenv("DATABASE_URL")
	if url == "" {
		url = "host=database dbname=vdelivery_db sslmode=disable port=5432 password=vdelivery_psw user=vdelivery_server"
	}

	log.Println("START SERVER")
	fmt.Println("START", port, "STARTfjdhjd")
	log.Println(app.Start(port, url))
}