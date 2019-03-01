package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type rgbJSON struct {
	Red   string
	Green string
	Blue  string
}

const (
	port = "3001"
)

var pigpio piGPIO

func formatString(format string, args ...interface{}) string {
	return fmt.Sprintf(format, args...)
}

func rgbHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var rgb rgbJSON
	err := decoder.Decode(&rgb)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(formatString("Red: %s, Green: %s, Blue: %s", rgb.Red, rgb.Green, rgb.Blue))
	pigpio.setRGB(rgb.Red, rgb.Green, rgb.Blue)
}

func main() {
	pigpio.redPin = "17"
	pigpio.greenPin = "22"
	pigpio.bluePin = "24"

	http.HandleFunc("/setRGB", rgbHandler)

	log.Println(formatString("staring server on port %s...", port))
	log.Fatal(http.ListenAndServe(formatString(":%s", port), nil))
}
