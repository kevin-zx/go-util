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
