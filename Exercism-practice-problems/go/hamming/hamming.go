package hamming

import (
	"fmt"
)

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, fmt.Errorf("Invalid comparison. Strings of unequal length.")
	}
	var diff int
	// aChars := strings.Split(a, "")
	// bChars := strings.Split(b, "")
	// for index, value := range aChars {
	// 	if value != bChars[index] {
	// 		diff += 1
	// 	}
	// }
	for index := range a {
		if a[index] != b[index] {
			// fmt.Printf(string(a[index]))
			diff += 1
		}
	}
	return diff, nil
}
