package _6

import (
	"AdventOfCode/Importer"
	"fmt"
	"strings"
)

const YearNumber = 2020
const DayNumber = 6
const FileName = "input.txt"

func GetResult1 () (string, error){
	input := Importer.ReadDayInput(YearNumber,DayNumber,FileName)
	groups := strings.Split(input, "\n\n")
	total := 0

	for _,g := range groups{
		group := strings.Split(g, "\n")
		total += GetGroupAnswerCount(group)
	}

	return fmt.Sprintf("%d", total), nil
}

func GetResult2 () (string, error){
	input := Importer.ReadDayInput(YearNumber,DayNumber,FileName)
	groups := strings.Split(input, "\n\n")
	total := 0

	for _,g := range groups{
		group := strings.Split(g, "\n")
		total += GetGroupFullAnswerCount(group)
	}

	return fmt.Sprintf("%d", total), nil
}

func GetGroupAnswerCount(group []string)int{
	questions := make(map[string]int)

	for _,person := range group{
		for _,q := range person{
			questions[string(q)]++
		}
	}

	return len(questions)
}

func GetGroupFullAnswerCount(group []string)int{
	questions := make(map[string]int)

	for _,person := range group{
		for _,q := range person{
			questions[string(q)]++
		}
	}

	allCorrect := 0
	for _, question := range questions{
		if question == len(group){
			allCorrect++
		}
	}

	return allCorrect
}
