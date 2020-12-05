package Importer

import (
	"fmt"
	"io/ioutil"
	"log"
)

func ReadDayInput(yearNumber, dayNumber int, fileName string) string{
	file, error := ioutil.ReadFile(fmt.Sprintf("Year/_%d/_%d/%s", yearNumber, dayNumber, fileName))
	if error != nil{
		log.Fatal(error)
	}
	return string(file)
}