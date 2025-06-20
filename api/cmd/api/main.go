package main

import (
	// Needs to be the first import since it loads the env vars
	_ "github.com/joho/godotenv/autoload" 

	"fmt"
	"net/http"
	"ios-app-reviews-viewer.com/m/internal/database"
)

func main() {
	database.Connect()

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "OK")
	})
	
	fmt.Println("Server starting on :8080")
	http.ListenAndServe(":8080", nil)
}
