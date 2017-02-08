package regexpUtil

import (
	"regexp"
)

//find match str from target line by regexp
func FindString(pattern string, line string) []string {
	re := regexp.MustCompile(pattern)
	resultStr := re.FindAllString(line, 1)
	return resultStr
}

func SplitString(s string, pattern string) []string {
	re, _ := regexp.Compile(pattern)
	return re.Split(s, -1)
}

func SplitString2Line(s string) []string {
	return SplitString(s, "\r\n|\r|\n")
}
