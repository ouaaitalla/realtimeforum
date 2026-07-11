package main

import (
	"fmt"
	"log"
	"net/http"

	"real-time-forum/backend/middleware"
	"real-time-forum/backend/routes"
	"real-time-forum/database"
)

func main() {
	if err := database.InitDB(); err != nil {
		log.Fatal(err)
	}
	defer database.CloseDB()

	// Start server...
	router := routes.SetupRoutes()

	handler := middleware.CORSMiddleware(router)

	log.Fatal(http.ListenAndServe(":8080", handler))
	fmt.Println("Server running on :8080")
	
}
