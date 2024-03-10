package main

import (
	"lsplanner-go/controllers"

	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	h := server.Default()

	h.GET("/ping", controllers.PingTest)
	h.GET("/users", controllers.HandleGetUsers)

	h.Spin()
}
