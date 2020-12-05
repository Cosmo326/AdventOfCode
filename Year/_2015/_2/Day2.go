package _2

import (
	"AdventOfCode/Importer"
	"fmt"
	"sort"
	"strings"
)

const YearNumber = 2015
const DayNumber = 2
const FileName = "input.txt"

func GetResult1 () (string, error) {
	input := Importer.ReadDayInput(YearNumber, DayNumber, FileName)
	inputArr := strings.Split(input, "\n")
	total := 0

	for _,v := range inputArr{
		var l, w, h int
		_, scanErr := fmt.Sscanf(v, "%dx%dx%d", &l, &w, &h)

		if scanErr != nil {
			panic(scanErr)
		}

		total += 2*l*w + 2*l*h + 2*h*w + GetMinArea(l,h,w)
	}

	return fmt.Sprintf("%d", total), nil
}

func GetMinArea(l,w,h int) int {
	if l*w <= l*h && l*w <= h*w {
		return l*w
	} else if l*h <= l*w && l*h <= h*w {
		return l*h
	} else {
		return h*w
	}
}

func GetResult2 () (string, error) {
	input := Importer.ReadDayInput(YearNumber, DayNumber, FileName)
	inputArr := strings.Split(input, "\n")
	var total int = 0

	for _,v := range inputArr{
		var dims = []int {0,0,0}
		_, scanErr := fmt.Sscanf(v, "%dx%dx%d", &dims[0], &dims[1], &dims[2])

		if scanErr != nil {
			panic(scanErr)
		}

		sort.Ints(dims)

		total += 2 * dims[0] + 2 * dims[1] + dims[0] * dims[1] * dims[2]
	}

	return fmt.Sprintf("%d", total), nil
}
