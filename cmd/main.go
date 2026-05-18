package main

import (
	"coffee-app-server/internal/db"
	"log"
)


func main(){
	dbConnection, err := db.ConnectDB()
	if err != nil {
		log.Fatalf("failed to connect db %v", err)
	}

	
}