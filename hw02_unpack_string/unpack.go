package hw02_unpack_string //nolint:golint,stylecheck

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(text string) (string, error) {
	var out strings.Builder
	var flag bool
	var nextNum bool
	var r rune
	out.Grow(len(text))
	for _, value := range text {
		switch {
		case value == '\\' && !flag:
			nextNum = false
			flag = true
		case flag:
			nextNum = false
			out.WriteRune(value)
			flag, r = false, value
		case unicode.IsNumber(value):
			if r != 0 && !nextNum {
				for n := 1; n < int(value-'0'); n++ {
					out.WriteRune(r)
				}
				nextNum = true
			} else {
				return "", errors.New("invalid string")
			}
		default:
			nextNum = false
			out.WriteRune(value)
			r = value
		}
	}
	return out.String(), nil
}
