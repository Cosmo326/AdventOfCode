package _5

import (
	"AdventOfCode/Importer"
	"fmt"
	"math"
	"sort"
	"strings"
)

const YearNumber = 2020
const DayNumber = 5
const FileName = "input.txt"

func GetResult1 () (string, error){
	input := Importer.ReadDayInput(YearNumber, DayNumber, FileName)
	inputArr := strings.Split(input, "\n")
	highestSeat := 0

	for _,v := range inputArr {
		var rowInst, colInst string
		_, scanErr := fmt.Sscanf(v, "%7s%3s", &rowInst, &colInst)

		if scanErr != nil {
			panic(scanErr)
		}
		row := TraverseTree(rowInst)
		col := TraverseTree(colInst)

		seat := row * 8 + col

		if seat > highestSeat {
			highestSeat = seat
		}
	}

	return fmt.Sprintf("%d", highestSeat), nil
}

func GetResult2 () (string, error){
	input := Importer.ReadDayInput(YearNumber, DayNumber, FileName)
	inputArr := strings.Split(input, "\n")
	var seats []int

	for _,v := range inputArr {
		var rowInst, colInst string
		_, scanErr := fmt.Sscanf(v, "%7s%3s", &rowInst, &colInst)

		if scanErr != nil {
			panic(scanErr)
		}
		row := TraverseTree(rowInst)
		col := TraverseTree(colInst)

		seat := row * 8 + col

		seats = append(seats, seat)
	}

	sort.Ints(seats)
	for s := 0; s < len(seats) - 1; s++{
		if seats[s] + 2 == seats[s + 1]{
			return fmt.Sprintf("%d", seats[s] + 1), nil
		}
	}

	return "Seat Doesn't Exist", nil
}

func TraverseTree(instructions string) int{
	position := 0
	half := int(math.Pow(2, float64(len(instructions)))) / 2

	for _, i := range instructions {
		if string(i) == "B" || string(i) == "R" {
			position += half
		}
		half /= 2
	}

	return position
}