package _3

import (
	"AdventOfCode/Importer"
	"fmt"
)

const YearNumber = 2015
const DayNumber = 3
const FileName = "input.txt"

func GetResult1 () (string, error){
	input := Importer.ReadDayInput(YearNumber,DayNumber,FileName)
	mapping := make(map[string]int)
	x := 0
	y := 0

	for _, c := range input{
		x,y = Move(string(c), x,y)
		mapping[fmt.Sprintf("%dx%d",x,y)]++
	}

	return fmt.Sprintf("%d", len(mapping)), nil
}

func GetResult2 () (string, error){
	input := Importer.ReadDayInput(YearNumber,DayNumber,FileName)
	mapping := make(map[string]int)
	santaX, roboX := 0,0
	santaY, roboY := 0,0

	mapping[fmt.Sprintf("%dx%d",santaX,santaY)]++
	mapping[fmt.Sprintf("%dx%d",roboX,roboY)]++
	for i, c := range input{
		if i % 2 == 0 {
			roboX, roboY = Move(string(c), roboX, roboY)
			mapping[fmt.Sprintf("%dx%d",roboX,roboY)]++
		} else {
			santaX, santaY = Move(string(c), santaX, santaY)
			mapping[fmt.Sprintf("%dx%d",santaX,santaY)]++
		}

	}

	return fmt.Sprintf("%d", len(mapping)), nil
}

func Move(instruction string, x,y int)(int, int){
	switch instruction{
	case ">":
		return x+1,y
	case "<":
		return x-1, y
	case "^":
		return x, y+1
	case "v":
		return x, y-1
	}
	return x,y
}