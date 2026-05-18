package main

import (
	"coffee-app-server/internal/db"
	"coffee-app-server/internal/order"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	dbConnection, err := db.ConnectDB()
	if err != nil {
		log.Fatalf("failed to connect db %v", err)
	}

	router := gin.Default()

	orderService := order.NewOrderService(dbConnection)
	orderHandler := order.NewOrderHandler(orderService)
	orderHandler.RegisterEndPoints(router)

	router.Run(":3000")
}
