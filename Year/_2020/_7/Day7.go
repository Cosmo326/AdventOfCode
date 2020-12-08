package _7

import (
	"AdventOfCode/Importer"
	"fmt"
	"strings"
)

const YearNumber = 2020
const DayNumber = 7
const FileName = "input.txt"

var Bags = make(map[string]map[string]int)

func PopulateBags(){
	if len(Bags) < 1 {
		input := Importer.ReadDayInput(YearNumber, DayNumber, FileName)
		rules := strings.Split(input, "\n")
		for _, rule := range rules {
			bag := strings.Split(rule, " bags contain ")
			bagContents := make(map[string]int)
			contents := strings.Split(bag[1], ", ")
			for _, content := range contents {
				var count int
				var description string
				_, scanErr := fmt.Sscanf(content, "%d %s bags", &count, &description)
				if scanErr != nil {
					panic(scanErr)
				}

				bagContents[description] = count
			}
			Bags[bag[0]] = bagContents
		}
	}
}

func GetResult1 () (string, error){


	return "Not Completed", nil
}

func GetResult2 () (string, error){
	return "Not Completed", nil
}
