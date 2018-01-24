package acronym //funcs: 1--practice coding w/ Go (package acronym).

import (
	"regexp"
	"strings"
)

//converts a phrase into an acronym of each word's first letter (func Abbreviate).
func Abbreviate(s string) interface{} {
	var acronym string

	words := regexp.MustCompile("[ -]").Split(s, -1) //building upon dotmnd's submission
	for _, word := range words {
		acronym += word[:1]
	}

	return strings.ToUpper(acronym)
}
