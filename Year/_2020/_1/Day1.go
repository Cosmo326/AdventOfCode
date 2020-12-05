package _1

import (
	"AdventOfCode/Importer"
	"AdventOfCode/Logger"
	"errors"
	"strconv"
	"strings"
)

const DayNumber = 1
const FileName = "input.txt"

func GetResult1 () (string, error) {
	inputString := Importer.ReadDayInput(DayNumber, FileName)
	arr := strings.Split(inputString, "\r\n")

	for i:=0; i < len(arr) - 1; i++{
		for j:=1; j < len(arr); j++{
			a, aErr := strconv.Atoi(arr[i])
			b, bErr := strconv.Atoi(arr[j])
			Logger.LogAllFatal([]error {aErr, bErr})
			if a + b == 2020 {
				return strconv.Itoa(a * b), nil
			}
		}
	}

	return "", errors.New("No Combo Exists")
}

func GetResult2() (string, error){
	inputString := Importer.ReadDayInput(DayNumber, FileName)
	arr := strings.Split(inputString, "\r\n")
	for i:=0; i < len(arr) - 2; i++{
		for j:=1; j < len(arr) - 1; j++{
			for k:=2; k < len(arr); k++ {
				a, aErr := strconv.Atoi(arr[i])
				b, bErr := strconv.Atoi(arr[j])
				c, cErr := strconv.Atoi(arr[k])
				Logger.LogAllFatal([] error{aErr, bErr, cErr})
				if a+b+c == 2020 {
					return strconv.Itoa(a * b * c), nil
				}
			}
		}
	}

	return "", errors.New("No Combo Exists")
}