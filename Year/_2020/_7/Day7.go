package _7

import (
	"AdventOfCode/Importer"
	"fmt"
	"strings"
)

const YearNumber = 2020
const DayNumber = 7
const FileName = "input.txt"

func PopulateBags() []Bag{
	Bags := make([]Bag, 0)
	input := Importer.ReadDayInput(YearNumber, DayNumber, FileName)
	rules := strings.Split(input, "\n")
	for _, rule := range rules {
		bagRule := strings.Split(rule, " bags contain ")
		bagContents := make([]BagContent,0)
		contents := strings.Split(bagRule[1], ", ")
		for _, content := range contents {
			var bagContent BagContent
			_, scanErr := fmt.Sscanf(content, "%d %s bags", &bagContent.Count, &bagContent.Name)
			if scanErr != nil {
				panic(scanErr)
			}

			bagContents = append(bagContents, bagContent)
		}
		bag := Bag{ Name: bagRule[0], Content: bagContents}
		Bags = append(Bags, bag)
	}
	return Bags
}

func GetResult1 () (string, error){
	Bags := PopulateBags()

	Where(Bags, func(bag Bag) bool {
		return Contains(bag.Content, func(bagContent BagContent) bool {
			return bagContent.Name == "shiny gold"
		})
	})


	return "Not Completed", nil
}

func GetResult2 () (string, error){
	return "Not Completed", nil
}
