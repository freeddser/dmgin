package handlers

import (
	"github.com/freeddser/dmgin/common"
	"github.com/freeddser/dmgin/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetDriverByID(c *gin.Context) {
	driverID := common.StringToInt(c.Param("id"))
	if driverID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Driver ID not found",
		})
		return
	}

	data, err := services.GetDriverByID(int64(driverID))
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSONP(http.StatusOK, gin.H{
		"msg":      "success",
		"response": data,
	})
}
