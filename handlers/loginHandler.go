package handlers

import (
	"github.com/freeddser/dmgin/request"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	var json request.LoginReq
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if json.Name != "gavin" || json.Password != "123456" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
}
