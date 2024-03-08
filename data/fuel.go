package data

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"time"
)

type Fuel struct {
	// ID       uint      `json:"id"`
	Name     string    `json:"name"`
	Company  string    `json:"company"`
	Price    int       `json:"price"`
	DateTime time.Time `json:"dateTime"`
}

// GetFuels reads the fuels.json file and returns a slice of Fuel structs
func GetFuels() ([]Fuel, error) {
	// Read the JSON file
	data, err := ioutil.ReadFile("fuels.json")
	if err != nil {
		log.Println("Error reading file:", err)
		return nil, err
	}

	// Unmarshal JSON data into a slice of Fuel structs
	var fuels []Fuel
	err = json.Unmarshal(data, &fuels)
	if err != nil {
		log.Println("Error unmarshalling JSON:", err)
		return nil, err
	}

	return fuels, nil
}
