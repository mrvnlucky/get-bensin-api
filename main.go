package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Fuel struct {
	ID        uint   `json:"id"`
	FuelName  string `json:"fuel_name"`
	FuelPrice string `json:"fuel_price"`
}

func main() {
	r := gin.Default()

	// Define your endpoint to handle the GET request
	r.GET("/fuels", func(c *gin.Context) {
		// Read fuel data from JSON file
		fuels, err := readFuelsFromJSON("fuels.json")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to read fuel data"})
			return
		}

		// Return fuels as JSON response
		c.JSON(http.StatusOK, fuels)
	})
	r.Run()
}

func readFuelsFromJSON(filename string) ([]Fuel, error) {
	// Open the JSON file
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Decode JSON data into a slice of Fuel structs
	var fuels []Fuel
	if err := json.NewDecoder(file).Decode(&fuels); err != nil {
		return nil, err
	}

	return fuels, nil
}
