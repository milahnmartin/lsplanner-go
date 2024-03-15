package main

import (
	"context"
	"log"
	"os"

	"lsplanner-go/config"
	"lsplanner-go/models"
	"lsplanner-go/repositories"
	"lsplanner-go/routes"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/redis/go-redis/v9"
)

func main() {
	optimalPort := envPortOr("8888")

	db, err := config.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := db.AutoMigrate(&models.User{}, &models.LoadsheddingArea{}, &models.Quota{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	redisAddress := os.Getenv("REDIS_URL")
	if redisAddress == "" {
		redisAddress = "localhost:6379"
	}
	redisPass := os.Getenv("REDIS_PASSW")
	if redisPass == "" {
		redisPass = ""
	}

	redisClient := redis.NewClient(&redis.Options{Addr: redisAddress, Password: redisPass, DB: 0})

	ctx := context.Background()
	_, err = redisClient.Ping(ctx).Result()
	if err != nil {
		log.Println("URL IS", os.Getenv("REDIS_URL"))
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	log.Println("Connected to Redis successfully.")

	h := server.New(server.WithHostPorts(optimalPort))

	userRepo := repositories.NewUserRepo(db)
	userGroup := h.Group("/api/v1/users")
	routes.UserRoutes(userGroup, userRepo)

	h.Spin()
}

func envPortOr(port string) string {
	if envPort := os.Getenv("PORT"); envPort != "" {
		return ":" + envPort
	}
	return ":" + port
}
