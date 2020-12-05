package _4

import (
	"AdventOfCode2020/Importer"
	"fmt"
	"log"
	"regexp"
	"strings"
)

const DayNumber = 4
const FileName = "input.txt"

func GetResult1 () (string, error){
	input := Importer.ReadDayInput(DayNumber, FileName)
	inputArr := strings.Split(input, "\r\n\r\n")
	valid := 0

	for _,v := range inputArr{
		regex, regexErr := regexp.Compile("byr|iyr|eyr|hgt|hcl|ecl|pid")
		if regexErr != nil{
			log.Fatal(regexErr)
		}
		fields := regex.FindAllString(v, -1)

		if len(fields) == 7 {
			valid++
		}
	}

	return fmt.Sprintf("%d", valid), nil
}

func GetResult2 () (string, error){
	input := Importer.ReadDayInput(DayNumber, FileName)
	inputArr := strings.Split(input, "\r\n\r\n")
	valid := 0

	for _,v := range inputArr{
		//regex, regexErr := regexp.Compile("(byr:[0-9]{4})|(iyr:[0-9]{4})|(eyr:[0-9]{4})|(hgt:[0-9]{2,3}[ci][mn])|(hcl:#[0-9a-f]{6})|(ecl:[a-z]{3})|(pid:[0-9]{9})")
		regex, regexErr := regexp.Compile("byr:\\w+|iyr:\\w+|eyr:\\w+|hgt:\\w+|hcl:#\\w+|ecl:\\w+|pid:\\w+")
		if regexErr != nil{
			log.Fatal(regexErr)
		}
		fields := regex.FindAllString(v, -1)

		if len(fields) == 7 {
			fieldsValid := true

			for _,v := range fields{
				if !TestField(v) {
					fieldsValid = false
					break
				}
			}

			if fieldsValid {
				valid++
			}
		}
	}

	// should be 194
	return fmt.Sprintf("%d", valid), nil
}

func TestField(field string) bool {
	items := strings.Split(field,":")
	switch items[0]{
	case "byr":
		return TestYear(items[1], 1920, 2002)
	case "iyr":
		return TestYear(items[1], 2010, 2020)
	case "eyr":
		return TestYear(items[1], 2020, 2030)
	case "hgt":
		return TestHeight(items[1])
	case "hcl":
		var hairColor string
		_, scanErr := fmt.Sscanf(items[1], "#%6s", &hairColor)
		if scanErr != nil{
			return false
		}
		return true
	case "ecl":
		return items[1] == "amb" || items[1] == "blu" || items[1] == "brn" || items[1] == "gry" || items[1] == "grn" || items[1] == "hzl" || items[1] == "oth"
	case "pid":
		return len(items[1]) == 9
	case "cid":
		println("You shouldn't be here")
		return true
	default:
		return false
	}
}

func TestYear(yearStr string, min, max int)bool{
	var year int
	_, scanErr := fmt.Sscanf(yearStr, "%4d", &year)

	if scanErr != nil{
		return false
	}
	return year >= min && year <= max
}

func TestHeight(height string)bool{
	var heightNumber int
	var heightMeasure string
	_, scanErr := fmt.Sscanf(height, "%d%2s", &heightNumber, &heightMeasure)
	if scanErr != nil{
		return false
	}

	if heightMeasure == "cm"{
		return heightNumber >= 150 && heightNumber <= 193
	} else if heightMeasure == "in" {
		return heightNumber >= 59 && heightNumber <= 76
	}

	return false
}