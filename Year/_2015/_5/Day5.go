package _5

import (
	"AdventOfCode/Importer"
	"fmt"
	"regexp"
	"strings"
)

const YearNumber = 2015
const DayNumber = 5
const FileName = "input.txt"

func GetResult1 () (string, error){
	input := Importer.ReadDayInput(YearNumber, DayNumber, FileName)
	inputArr := strings.Split(input, "\n")
	naughty := 0
	nice := 0

	for _,v := range inputArr{
		if CheckBadStrings(1, v) {
			//println(v)
			naughty++
		} else if CheckDoubles(1, v) && CheckVowels(3, v) {
			//println(v)
			nice++
		} else {
			println(v)
			naughty++
		}
	}

	return fmt.Sprintf("%d", nice), nil
}

func GetResult2 () (string, error){
	return "Not Completed", nil
}

func CheckVowels(count int, line string) bool{
	regex, _ := regexp.Compile("[aeiou]{1}")
	matches := regex.FindAllString(line, -1)
	return len(matches) >= count
}

func CheckDoubles(count int, line string) bool{
	regex, _ := regexp.Compile("aa|bb|cc|dd|ee|ff|gg|hh|ii|jj|kk|ll|mm|nn|oo|pp|qq|rr|ss|tt|uu|vv|ww|xx|yy|zz")
	matches := regex.FindAllString(line, -1)
	return len(matches) >= count
}

func CheckBadStrings(count int, line string) bool{
	regex, _ := regexp.Compile("ab|cd|pq|xy")
	matches := regex.FindAllString(line, -1)
	return len(matches) >= count
}
