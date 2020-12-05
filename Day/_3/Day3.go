package _3

import (
	"AdventOfCode2020/Importer"
	"fmt"
	"strings"
)

const DayNumber = 3
const FileName = "input.txt"

func GetResult1 () (string, error){
	inputString := Importer.ReadDayInput(DayNumber, FileName)
	hillMap := strings.Split(inputString, "\r\n")

	treesEncountered := GetTreesEncountered(hillMap, 3, 1)

	return fmt.Sprintf("%d", treesEncountered), nil
}

func GetResult2 () (string, error){
	inputString := Importer.ReadDayInput(DayNumber, FileName)
	hillMap := strings.Split(inputString, "\r\n")

	treesEncountered := GetTreesEncountered(hillMap, 1, 1)
	treesEncountered *= GetTreesEncountered(hillMap, 3, 1)
	treesEncountered *= GetTreesEncountered(hillMap, 5, 1)
	treesEncountered *= GetTreesEncountered(hillMap, 7, 1)
	treesEncountered *= GetTreesEncountered(hillMap, 1, 2)

	return fmt.Sprintf("%d", treesEncountered), nil
}

func GetTreesEncountered(hill []string, left, drop int) int{
	treesEncountered := 0
	currentRight := 0

	for i := 0; i < len(hill); i += drop{
		nextStop := (hill[i])[currentRight: currentRight+1]

		if nextStop == "#"{
			treesEncountered++
		}

		currentRight = (currentRight + left) % len(hill[i])
	}

	return treesEncountered
}