package api

import (
	"get-bensin/data"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterFuelRoutes(router *gin.Engine) {
	router.GET("/fuels", getFuelsHandler)
}

func getFuelsHandler(c *gin.Context) {
	fuels, err := data.GetFuels()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get fuels"})
		return
	}

	c.JSON(http.StatusOK, fuels)
}
