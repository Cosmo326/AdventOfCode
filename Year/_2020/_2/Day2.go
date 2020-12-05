package _2

import (
	"AdventOfCode/Importer"
	"fmt"
	"log"
	"regexp"
	"strings"
)

const YearNumber = 2020
const DayNumber = 2
const FileName = "input.txt"

func GetResult1 () (string, error) {
	inputString := Importer.ReadDayInput(YearNumber, DayNumber, FileName)
	inputArr := strings.Split(inputString, "\r\n")

	successCount := 0

	for _, v := range inputArr {
		min, max, checkLetter, pwrd := ParseLine(v)

		regex, regexError := regexp.Compile(fmt.Sprintf("%s",checkLetter))
		if regexError != nil {
			log.Fatal(regexError)
		}
		matches := regex.FindAllString(pwrd, -1)

		if matches != nil && len(matches) >= min && len(matches) <= max{
			successCount++
		}
	}

	return fmt.Sprintf("%d",successCount), nil
}

func GetResult2 () (string, error) {
	inputString := Importer.ReadDayInput(YearNumber, DayNumber, FileName)
	inputArr := strings.Split(inputString, "\r\n")

	successCount := 0

	for _, v := range inputArr {
		firstEnd, secondEnd, checkLetter, pwrd := ParseLine(v)

		first := pwrd[(firstEnd-1):firstEnd]
		second := pwrd[(secondEnd-1):secondEnd]

		if (first == checkLetter || second == checkLetter) && !(first == checkLetter && second == checkLetter){
			successCount++
		}
	}

	return fmt.Sprintf("%d",successCount), nil
}

func ParseLine(line string)(int, int, string, string) {
	var first, second int
	var checkLetter, password string

	n, scanErr := fmt.Sscanf(line, "%d-%d %1s: %s", &first, &second, &checkLetter, &password)
	if scanErr != nil{
		println(line)
		log.Fatal(scanErr)
	}

	if n != 4{
		panic(fmt.Sprintf("Line not formatted correctly: %s", line))
	}

	return first, second, checkLetter, password

}
