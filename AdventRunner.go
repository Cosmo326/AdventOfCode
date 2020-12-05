package main

import (
	"AdventOfCode/Year"
	"flag"
	"fmt"
	"log"
	"time"
)

func main() {

	year := flag.Int("Year", time.Now().Year(), "Choose which year to run. Default is current year")
	day := flag.Int("_2020", time.Now().Day(), "Choose which day to run. Default is current day")
	all := flag.Bool("All", false, "Run all days")
	allYear := flag.Bool("AllYear", false, "Run all days for a specific year.  If year flag is not specified current year will be run")

	flag.Parse()

	if *all {
		for y := 2015; y <= *year; y++ {
			for d := 1; d < 26; d++ {
				RunDay(y, d)
			}
		}
	} else if *allYear {
		for d := 1; d < 26; d++ {
			RunDay(*year, d)
		}
	} else if time.Now().Month() == 12 && time.Now().Day() < 26 {
		RunDay(*year, *day)
	} else {
		println("Advent of Code is not currently running.  Check the Help menu to see how to rerun any and all days")
	}
}

func RunDay(yearNumber, dayNumber int){
	result, dayRunnerError := Year.RunYear(yearNumber, dayNumber)

	if dayRunnerError != nil{
		log.Fatal(dayRunnerError)
	}

	println()
	println(fmt.Sprintf("12/%d/%d", dayNumber, yearNumber))
	println(fmt.Sprintf("  Answer 1 is: %s", result[0]))
	println(fmt.Sprintf("  Answer 2 is: %s", result[1]))
	println()
}




