package main

import (
	"log"
	"net/http"
	"warehouse_inventory/config"
	"warehouse_inventory/database"
	"warehouse_inventory/routes"
)

func main() {
	// Load configuration
	if err := config.LoadConfig(); err != nil {
		log.Fatal("Error loading config:", err)
	}

	// Initialize database
	db, err := database.InitDB()
	if err != nil {
		log.Fatal("Error initializing database:", err)
	}
	defer db.Close()

	router := routes.SetupRouter()

	log.Printf("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
