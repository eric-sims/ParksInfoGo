package main

import (
	"net/http"
	"parksInfoGo/controllers"

	"github.com/gin-gonic/gin"

	_ "github.com/dimiro1/banner/autoload"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/disney-queues", controllers.GetDisneyQueues)
	r.GET("/ca-queues", controllers.GetCaliforniaAdventureQueues)

	err := r.Run()
	if err != nil {
		panic(err)
	}
}
