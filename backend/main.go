package main

import (
	"log"
	"net/http"

	"real-time-forum/backend/helpers"
	"real-time-forum/backend/middleware"
	"real-time-forum/backend/routes"
	"real-time-forum/backend/websocket"
	"real-time-forum/database"
)

func main() {
	if err := database.InitDB(); err != nil {
		log.Fatal(err)
	}
	defer database.CloseDB()
	defer helpers.LoginRateLimiter.Stop()

	// Start server...
	router := routes.SetupRoutes()

	handler := middleware.CORSMiddleware(router)

	go websocket.HubInstance.Run()

	log.Fatal(http.ListenAndServe(":8080", handler))
}
