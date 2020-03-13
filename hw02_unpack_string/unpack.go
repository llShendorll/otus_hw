package hw02_unpack_string //nolint:golint,stylecheck

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(text string) (string, error) {
	var out strings.Builder
	out.Grow(len(text))
	var hasBackslash bool
	var nextNum bool
	var r rune
	for _, value := range text {
		switch {
		case value == '\\' && !hasBackslash:
			nextNum = false
			hasBackslash = true
		case hasBackslash:
			nextNum = false
			out.WriteRune(value)
			hasBackslash, r = false, value
		case unicode.IsDigit(value):
			if r != 0 && !nextNum {
				out.WriteString(strings.Repeat(string(r), int(value-'1')))
				nextNum = true
			} else {
				return "", ErrInvalidString
			}
		default:
			nextNum = false
			out.WriteRune(value)
			r = value
		}
	}
	return out.String(), nil
}
