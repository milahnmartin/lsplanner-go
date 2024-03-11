package main

import (
	"lsplanner-go/controllers"
	"os"

	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {

	optimalPort := envPortOr("8888")

	h := server.New(server.WithHostPorts(optimalPort))

	h.GET("/ping", controllers.PingTest)
	h.GET("/users", controllers.HandleGetUsers)

	h.Spin()
}

func envPortOr(port string) string {
	// If `PORT` variable in environment exists, return it
	if envPort := os.Getenv("PORT"); envPort != "" {
		return ":" + envPort
	}
	// Otherwise, return the value of `port` variable from function argument
	return ":" + port
}
