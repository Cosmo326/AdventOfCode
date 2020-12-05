package _2019

import (
	"AdventOfCode/Logger"
	"AdventOfCode/Year/_2019/_1"
	"AdventOfCode/Year/_2019/_10"
	"AdventOfCode/Year/_2019/_11"
	"AdventOfCode/Year/_2019/_12"
	"AdventOfCode/Year/_2019/_13"
	"AdventOfCode/Year/_2019/_14"
	"AdventOfCode/Year/_2019/_15"
	"AdventOfCode/Year/_2019/_16"
	"AdventOfCode/Year/_2019/_17"
	"AdventOfCode/Year/_2019/_18"
	"AdventOfCode/Year/_2019/_19"
	"AdventOfCode/Year/_2019/_2"
	"AdventOfCode/Year/_2019/_20"
	"AdventOfCode/Year/_2019/_21"
	"AdventOfCode/Year/_2019/_22"
	"AdventOfCode/Year/_2019/_23"
	"AdventOfCode/Year/_2019/_24"
	"AdventOfCode/Year/_2019/_25"
	"AdventOfCode/Year/_2019/_3"
	"AdventOfCode/Year/_2019/_4"
	"AdventOfCode/Year/_2019/_5"
	"AdventOfCode/Year/_2019/_6"
	"AdventOfCode/Year/_2019/_7"
	"AdventOfCode/Year/_2019/_8"
	"AdventOfCode/Year/_2019/_9"
)

type GetResult func () (string, error)

func Run(dayNumber int) ([]string, error){
	var runRes1 GetResult
	var runRes2 GetResult

	switch dayNumber {
	case 1:
		runRes1 = _1.GetResult1
		runRes2 = _1.GetResult2
	case 2:
		runRes1 = _2.GetResult1
		runRes2 = _2.GetResult2
	case 3:
		runRes1 = _3.GetResult1
		runRes2 = _3.GetResult2
	case 4:
		runRes1 = _4.GetResult1
		runRes2 = _4.GetResult2
	case 5:
		runRes1 = _5.GetResult1
		runRes2 = _5.GetResult2
	case 6:
		runRes1 = _6.GetResult1
		runRes2 = _6.GetResult2
	case 7:
		runRes1 = _7.GetResult1
		runRes2 = _7.GetResult2
	case 8:
		runRes1 = _8.GetResult1
		runRes2 = _8.GetResult2
	case 9:
		runRes1 = _9.GetResult1
		runRes2 = _9.GetResult2
	case 10:
		runRes1 = _10.GetResult1
		runRes2 = _10.GetResult2
	case 11:
		runRes1 = _11.GetResult1
		runRes2 = _11.GetResult2
	case 12:
		runRes1 = _12.GetResult1
		runRes2 = _12.GetResult2
	case 13:
		runRes1 = _13.GetResult1
		runRes2 = _13.GetResult2
	case 14:
		runRes1 = _14.GetResult1
		runRes2 = _14.GetResult2
	case 15:
		runRes1 = _15.GetResult1
		runRes2 = _15.GetResult2
	case 16:
		runRes1 = _16.GetResult1
		runRes2 = _16.GetResult2
	case 17:
		runRes1 = _17.GetResult1
		runRes2 = _17.GetResult2
	case 18:
		runRes1 = _18.GetResult1
		runRes2 = _18.GetResult2
	case 19:
		runRes1 = _19.GetResult1
		runRes2 = _19.GetResult2
	case 20:
		runRes1 = _20.GetResult1
		runRes2 = _20.GetResult2
	case 21:
		runRes1 = _21.GetResult1
		runRes2 = _21.GetResult2
	case 22:
		runRes1 = _22.GetResult1
		runRes2 = _22.GetResult2
	case 23:
		runRes1 = _23.GetResult1
		runRes2 = _23.GetResult2
	case 24:
		runRes1 = _24.GetResult1
		runRes2 = _24.GetResult2
	case 25:
		runRes1 = _25.GetResult1
		runRes2 = _25.GetResult2
	default:
		return []string{"Not Implemented", "Not Implemented"}, nil
	}

	result1, result1Error := runRes1()
	result2, result2Error := runRes2()

	Logger.LogAllFatal([] error {result1Error, result2Error})

	return []string{result1, result2}, nil
}
