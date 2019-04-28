package stringUtil

import "strings"

func Clear(s string) string {
	s = strings.TrimSpace(s)
	s = strings.Trim(s, " ")
	s = strings.Replace(s, "​", "", -1)
	s = strings.Replace(s, "\n", "", -1)
	return s
}
