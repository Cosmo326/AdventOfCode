package Year

import (
	"AdventOfCode/Year/_2020"
	"AdventOfCode/Year/_2019"
	"AdventOfCode/Year/_2018"
	"AdventOfCode/Year/_2017"
	"AdventOfCode/Year/_2016"
	"AdventOfCode/Year/_2015"
)

func RunYear(year, day int) ([]string, error) {
	switch year{
	case 2020:
		return _2020.Run(day)
	case 2019:
		return _2019.Run(day)
	case 2018:
		return _2018.Run(day)
	case 2017:
		return _2017.Run(day)
	case 2016:
		return _2016.Run(day)
	case 2015:
		return _2015.Run(day)
	default:
		return []string{"Not Implemented","Not Implemented"}, nil
	}

}