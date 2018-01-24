package twofer //funcs: 1--practice w/ Go (package twofer).

import "bytes"

func ShareWith(name string) string { //adds name into a set phrase (func ShareWith).
	if name == "" {
		return "One for you, one for me."
	} else {
		//return fmt.Sprintf("One for %s, one for me.", s)
		var output bytes.Buffer
		output.WriteString("One for ")
		output.WriteString(name)
		output.WriteString(", one for me.")
		return output.String()
	}
}
