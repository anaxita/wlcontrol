package entity

import (
	"unicode/utf8"
)

func validateCharLen(text string, min, max int) bool {
	n := utf8.RuneCountInString(text)

	return n >= min && n <= max
}
