package _4

import (
	"crypto/md5"
	"fmt"
	"regexp"
)

const input = "bgvyzdsv"

func GetResult1 () (string, error){
	regex, _ := regexp.Compile("^00000.+")

	return fmt.Sprintf("%d", FindCompletionInt(regex)), nil
}

func GetResult2 () (string, error){
	regex, _ := regexp.Compile("^000000.+")

	return fmt.Sprintf("%d", FindCompletionInt(regex)), nil
}

func FindCompletionInt(regex *regexp.Regexp) int{
	i:=1
	for {
		data := []byte(fmt.Sprintf("%s%d", input, i))
		hex := fmt.Sprintf("%x", md5.Sum(data))
		if regex.MatchString(hex) {
			return i
		}
		i++
	}
}