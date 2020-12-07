package _6

import (
	"AdventOfCode/Importer"
	"fmt"
	"strings"
)

const YearNumber = 2015
const DayNumber = 6
const FileName = "input.txt"

var LightsOn = make(map[string]bool)
var LightsBrightness = make(map[string]int)

func PopulateLights(){
	input := Importer.ReadDayInput(YearNumber, DayNumber, FileName)
	instructions := strings.Split(input, "\n")

	for _,i := range instructions {
		var startX, startY, endX, endY int

		action := "toggle"
		format := "toggle %d,%d through %d,%d"
		if strings.Contains(i, "turn on") {
			action = "turn on"
			format = "turn on %d,%d through %d,%d"
		} else if strings.Contains(i, "turn off") {
			action = "turn off"
			format = "turn off %d,%d through %d,%d"
		}

		_, scanErr := fmt.Sscanf(i, format, &startX, &startY, &endX, &endY)
		if scanErr != nil {
			panic(scanErr)
		}

		switch action {
		case "toggle":
			Toggle(startX, endX, startY, endY)
		case "turn on":
			TurnOn(startX, endX, startY, endY)
		case "turn off":
			TurnOff(startX, endX, startY, endY)
		}
	}
}

func GetResult1 () (string, error){
	if len(LightsOn) == 0 {
		PopulateLights()
	}

	return fmt.Sprintf("%d", len(LightsOn)), nil
}

func GetResult2 () (string, error){
	brightness := 0

	if len(LightsBrightness) == 0 {
		PopulateLights()
	}

	for _, b := range LightsBrightness{
		brightness += b
	}

	return fmt.Sprintf("%d", brightness), nil
}

func Toggle(startX, endX, startY, endY int){
	for x := startX; x <= endX; x++ {
		for y := startY; y <= endY; y++ {
			if LightsOn[fmt.Sprintf("%dx%d", x, y)] {
				delete(LightsOn, fmt.Sprintf("%dx%d", x, y))
			} else {
				LightsOn[fmt.Sprintf("%dx%d", x, y)] = true
			}
			LightsBrightness[fmt.Sprintf("%dx%d", x, y)] += 2
		}
	}
}

func TurnOn(startX, endX, startY, endY int){
	for x := startX; x <= endX; x++ {
		for y := startY; y <= endY; y++ {
			LightsOn[fmt.Sprintf("%dx%d", x, y)] = true
			LightsBrightness[fmt.Sprintf("%dx%d", x, y)]++
		}
	}
}

func TurnOff(startX, endX, startY, endY int){
	for x := startX; x <= endX; x++ {
		for y := startY; y <= endY; y++ {
			delete(LightsOn, fmt.Sprintf("%dx%d", x, y))
			LightsBrightness[fmt.Sprintf("%dx%d", x, y)]--
			if LightsBrightness[fmt.Sprintf("%dx%d", x, y)] < 0 {
				delete(LightsBrightness, fmt.Sprintf("%dx%d", x, y))
			}
		}
	}
}
