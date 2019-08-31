package router

import (
	"github.com/freeddser/dmgin/handlers"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	// ping
	r.GET("/ping", handlers.Ping)

	// login
	r.POST("/login", handlers.Login)

	// Driver
	r.GET("/driver/:id", handlers.GetDriverByID)

	return r
}
