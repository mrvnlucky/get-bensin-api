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

func WriteJSON(data *[]data.Fuel) {
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		fmt.Println("Unable to create JSON file")
		return
	}
	_ = os.WriteFile("fuels.json", file, 0644)
}

func ToIDR(str string) int {
	if str == "-" || str == "N/A" {
		return 0
	}
	trimmed := strings.TrimPrefix(str, "IDR")
	trimmed = strings.TrimPrefix(trimmed, "Rp")
	trimmed = strings.ReplaceAll(trimmed, ",", "")
	trimmed = strings.TrimSpace(trimmed)
	num, err := strconv.Atoi(trimmed)
	if err != nil {
		log.Fatalln("Error:", err)
		return -1
	}
	return num
}
