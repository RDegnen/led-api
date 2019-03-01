package main

import (
	"log"
	"os/exec"
)

type piGPIO struct {
	redPin   string
	greenPin string
	bluePin  string
}

func (pigs *piGPIO) setRGB(redValue string, greenValue string, blueValue string) {
	setRed := exec.Command("pigs", "p", pigs.redPin, redValue)
	setGreen := exec.Command("pigs", "p", pigs.greenPin, greenValue)
	setBlue := exec.Command("pigs", "p", pigs.bluePin, blueValue)

	if err := setRed.Run(); err != nil {
		log.Println(err.Error())
	}
	if err := setGreen.Run(); err != nil {
		log.Println(err.Error())
	}
	if err := setBlue.Run(); err != nil {
		log.Println(err.Error())
	}
}
