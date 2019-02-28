package main

import (
	"os"
)

func main() {
	red := os.Args[1]
	green := os.Args[2]
	blue := os.Args[3]

	pigpio := piGPIO{redPin: "17", greenPin: "22", bluePin: "24"}
	pigpio.setRGB(red, green, blue)
}
