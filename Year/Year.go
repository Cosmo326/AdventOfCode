package Year

import (
	"AdventOfCode/Year/_2020"
)

func RunYear(year, day int) ([]string, error) {
	switch year{
	case 2020:
		return _2020.Run(day)
	default:
		return []string{"Not Implemented","Not Implemented"}, nil
	}

}