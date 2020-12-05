package _1

import (
	"AdventOfCode/Importer"
	"AdventOfCode/Logger"
	"fmt"
	"regexp"
)

const YearNumber = 2015
const DayNumber = 1
const FileName = "input.txt"

func GetResult1 () (string, error) {
	input := Importer.ReadDayInput(YearNumber, DayNumber, FileName)
	upPattern, upPatternErr := regexp.Compile("\\({1}")
	downPattern, downPatternErr := regexp.Compile("\\){1}")
	Logger.LogAllFatal([]error {upPatternErr, downPatternErr})

	up := upPattern.FindAllString(input, -1)
	down := downPattern.FindAllString(input, -1)

	return fmt.Sprintf("%d", len(up) - len(down)), nil
}

func GetResult2() (string, error){
	input := Importer.ReadDayInput(YearNumber, DayNumber, FileName)
	currentLevel := 0

	for i,c := range input{
		if string(c) == "(" {
			currentLevel++
		} else if string(c) == ")" {
			currentLevel--
		}

		if currentLevel < 0{
			return fmt.Sprintf("%d", i+1), nil
		}

	}

	return "Never went to basement", nil
}