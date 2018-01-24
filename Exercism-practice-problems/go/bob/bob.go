// Package bob: one random function, for practicing Go.
package bob

import "strings"

// func Hey: Bob will respond formulaically.
func Hey(remark string) string {
	//OR: switch { case remark == /* */: /* */}
	if strings.HasSuffix(remark, "?") &&
		remark == strings.ToUpper(remark) &&
		remark != strings.ToLower(remark) {
		return "Calm down, I know what I'm doing!"
	} else if strings.HasSuffix(remark, "?") {
		return "Sure."
	} else if remark == strings.ToUpper(remark) && remark != strings.ToLower(remark) {
		return "Whoa, chill out!"
	} else if len(remark) == 0 || strings.Count(remark, " ") == len(remark) {
		return "Fine. Be that way!"
	} else {
		return "Whatever."
	}
}
