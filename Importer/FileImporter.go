package Importer

import (
	"fmt"
	"io/ioutil"
	"log"
)

func ReadDayInput(dayNumber int, fileName string) string{
	file, error := ioutil.ReadFile(fmt.Sprintf("_2020/_%d/%s", dayNumber, fileName))
	if error != nil{
		log.Fatal(error)
	}
	return string(file)
}