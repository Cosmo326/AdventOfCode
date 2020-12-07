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
			naughty++
		} else if CheckDoubles(1, v) && CheckVowels(3, v) {
			nice++
		} else {
			naughty++
		}
	}

	return fmt.Sprintf("%d", nice), nil
}

func GetResult2 () (string, error){
	input := Importer.ReadDayInput(YearNumber, DayNumber, FileName)
	inputArr := strings.Split(input, "\n")
	naughty := 0
	nice := 0

	for _,v := range inputArr{
		if CheckDuplicatePairs(1, v) && CheckSingleSeparatedDoubles(1, v) {
			nice++
		} else {
			naughty++
		}
	}

	return fmt.Sprintf("%d", nice), nil
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

func CheckSingleSeparatedDoubles(count int, line string) bool{
	regex, _ := regexp.Compile("a.a|b.b|c.c|d.d|e.e|f.f|g.g|h.h|i.i|j.j|k.k|l.l|m.m|n.n|o.o|p.p|q.q|r.r|s.s|t.t|u.u|v.v|w.w|x.x|y.y|z.z")
	matches := regex.FindAllString(line, -1)
	return len(matches) >= count
}

func CheckDuplicatePairs(count int, line string) bool{
	subs := make(map[string]int)
	triples := CountTriples(line)

	for i:=0; i < len(line) - 1; i++{
		subs[line[i:i+2]]++
	}

	return len(line) - 1 - (len(subs) + triples) >= count
}

func CountTriples(line string) int{
	regex, _ := regexp.Compile("aaa|bbb|ccc|ddd|eee|fff|ggg|hhh|iii|jjj|kkk|lll|mmm|nnn|ooo|ppp|qqq|rrr|sss|ttt|uuu|vvv|www|xxx|yyy|zzz")
	matches := regex.FindAllString(line, -1)
	return len(matches)
}

func CheckBadStrings(count int, line string) bool{
	regex, _ := regexp.Compile("ab|cd|pq|xy")
	matches := regex.FindAllString(line, -1)
	return len(matches) >= count
}
