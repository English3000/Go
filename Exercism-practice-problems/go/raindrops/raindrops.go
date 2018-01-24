package raindrops

import (
	"bytes"
	"strconv"
)

func Convert(num int) string {
	var factorSounds bytes.Buffer //for better concat performance
	if num%3 == 0 {
		factorSounds.WriteString("Pling")
	}
	if num%5 == 0 {
		factorSounds.WriteString("Plang")
	}
	if num%7 == 0 {
		factorSounds.WriteString("Plong")
	}
	if len(factorSounds.String()) == 0 {
		return strconv.Itoa(num) //string(num) returns differently
	}
	return factorSounds.String()
}
