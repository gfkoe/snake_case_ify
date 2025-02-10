package convert

import (
	"unicode"
)

func Convert(str string) string {
	s := ""

	for i := 0; i < len(str); i++ {
		var c rune = rune(str[i])
		if unicode.IsUpper(c) {
			if i != 0 {
				s += "_" + string(unicode.ToLower(c))
			} else {
				s += string(unicode.ToLower(c))
			}
		} else {
			s += string(c)
		}
	}

	return s
}
