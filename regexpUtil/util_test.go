package regexpUtil

import (
	"fmt"
	"testing"
)

func testFindString(t *testing.T) {
	var tests = []struct {
		pattern string
		line    string
	}{
		{"123abc", "\\d"},
	}
	for _, test := range tests {
		if len(FindString(test.pattern, test.line)) == 0 {
			t.Error("regexUtil can't find str")
		}
	}
}
