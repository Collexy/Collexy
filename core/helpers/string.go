package helpers

import (
	"unicode"
	"unicode/utf8"
)

func IsFirstLower(s string) bool {

	r, _ := utf8.DecodeRuneInString(s)
	return bool(unicode.IsLower(r))
}

func IsFirstUpper(s string) bool {

	r, _ := utf8.DecodeRuneInString(s)
	return bool(unicode.IsUpper(r))
}

func IsFirstNumber(s string) bool {

	r, _ := utf8.DecodeRuneInString(s)
	return bool(unicode.IsNumber(r))
}
