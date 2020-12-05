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
	} else {
		RunDay(*day)
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




