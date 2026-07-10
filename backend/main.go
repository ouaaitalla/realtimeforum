package main

import (
	"log"
	"real-time-forum/database"
)

func main() {
	if err := database.InitDB(); err != nil {
		log.Fatal(err)
	}
	defer database.CloseDB()
	
	// Start server...
}