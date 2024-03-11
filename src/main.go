package main

import (
	"log"
	"os"

	"lsplanner-go/config"
	"lsplanner-go/controllers"
	"lsplanner-go/models"

	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	optimalPort := envPortOr("8888")

	db, err := config.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Running database migrations
	if err := db.AutoMigrate(&models.Quota{}, &models.User{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Initialize the controllers with the database instance
	controllers.Initialize(db)

	h := server.New(server.WithHostPorts(optimalPort))

	h.GET("/users/:id", controllers.GetUser)
	h.GET("/users", controllers.GetAllUsers) // Assuming you have a GetAllUsers function in your controllers
	h.POST("/users", controllers.CreateUser)
	h.PUT("/users/:id", controllers.UpdateUser)    // Assuming you have UpdateUser function in your controllers
	h.DELETE("/users/:id", controllers.DeleteUser) // Assuming you have DeleteUser function in your controllers

	h.Spin()
}

func envPortOr(port string) string {
	if envPort := os.Getenv("PORT"); envPort != "" {
		return ":" + envPort
	}
	return ":" + port
}
