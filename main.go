package main

import (
	"fmt"
	"os"
	"os/exec"
)

func runCommand(color string, value string) {
	command := exec.Command("pigs", "p", color, value)
	err := command.Run()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func main() {
	color := os.Args[1]
	rgbValue := os.Args[2]
	colorPins := map[string]string{
		"red":   "17",
		"green": "22",
		"blue":  "24",
	}

	runCommand(colorPins[color], rgbValue)
}
