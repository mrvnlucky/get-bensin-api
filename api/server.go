package api

import (
	"github.com/gin-gonic/gin"
)

// SetupRouter sets up the API routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	// GET request to /fuels endpoint
	r.GET("/fuels", getFuelsHandler)

	return r
}
