package convert

import (
	"unicode"
)

func Convert(str string) string {
	s := ""

	for i := 0; i < len(str); i++ {
		var c rune = rune(str[i])
		// for kebab
		if string(c) == "-" {
			s += "_"
		} else if unicode.IsUpper(c) { // for camel and pascal
			if i != 0 {
				s += "_" + string(unicode.ToLower(c))
			} else {
				s += string(unicode.ToLower(c))
			}
		} else { // add the char/rune
			s += string(c)
		}
	}

	return s
}
