package main

import (
	"AdventOfCode2020/Day"
	"flag"
	"fmt"
	"log"
	"time"
)

func main() {

	day := flag.Int("Day", time.Now().Day(), "Choose which day to run")
	all := flag.Bool("All", false, "Run All Days")

	flag.Parse()

	if *all {
		for i := 1; i < 26; i++ {
			RunDay(i)
		}
	} else if time.Now().Month() == 12 && time.Now().Year() == 2020 && time.Now().Day() < 26 {
		RunDay(*day)
	} else {
		println("2020 Advent of Code has completed.  Check the Help menu to se how to rerun any and all days")
	}
}

func RunDay(dayNumber int){
	result, dayRunnerError := Day.Run(dayNumber)

	if dayRunnerError != nil{
		log.Fatal(dayRunnerError)
	}

	println()
	println(fmt.Sprintf("Day %d", dayNumber))
	println(fmt.Sprintf("    Answer 1 is: %s", result[0]))
	println(fmt.Sprintf("    Answer 2 is: %s", result[1]))
	println()
}




