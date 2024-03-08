package util

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"get-bensin/data"
)

// WriteJSON writes the given fuel data to a JSON file
func WriteJSON(data *[]data.Fuel) {
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		fmt.Println("Unable to create JSON file")
		return
	}
	_ = os.WriteFile("fuels.json", file, 0644)
}

// ToIDR converts a string representation of Indonesian Rupiah to an integer
// value. It returns 0 if the input is either "-" or "N/A".
func ToIDR(str string) int {
	if str == "-" || str == "N/A" {
		return 0
	}
	// trim leading "IDR" and "Rp"
	trimmed := strings.TrimPrefix(str, "IDR")
	trimmed = strings.TrimPrefix(trimmed, "Rp")
	// replace all commas with empty strings
	trimmed = strings.ReplaceAll(trimmed, ",", "")
	// trim leading and trailing whitespace
	trimmed = strings.TrimSpace(trimmed)
	// convert to an integer
	num, err := strconv.Atoi(trimmed)
	if err != nil {
		log.Fatalln("Error:", err)
		return -1
	}
	return num
}
