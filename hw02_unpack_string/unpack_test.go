package hw02_unpack_string //nolint:golint,stylecheck

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

type test struct {
	input    string
	expected string
	err      error
}

func TestUnpackOk(t *testing.T) {
	for _, tst := range [...]test{
		{
			input:    "a4bc2d5e",
			expected: "aaaabccddddde",
		},
		{
			input:    "abcd",
			expected: "abcd",
		},
	} {
		result, _ := Unpack(tst.input)
		require.Equal(t, tst.expected, result)
	}
}

func TestUnpackErr(t *testing.T) {
	items := []string{
		"35",
		"",
		"dd22",
		"1\\dd22",
	}
	for _, value := range items {
		result, err := Unpack(value)
		if err == nil && result == "" {
			errors.New("invalid string ")
		} else {
			require.Equal(t, ErrInvalidString, err)
		}

	}
}

func TestUnpackWithEscape(t *testing.T) {
	//t.Skip() // Remove if task with asterisk completed

	for _, tst := range [...]test{
		{
			input:    `qwe\4\5`,
			expected: `qwe45`,
		},
		{
			input:    `qwe\45`,
			expected: `qwe44444`,
		},
		{
			input:    `qwe\\5`,
			expected: `qwe\\\\\`,
		},
		{
			input:    `qwe\\\3`,
			expected: `qwe\3`,
		},
	} {
		result, err := Unpack(tst.input)
		require.Equal(t, tst.err, err)
		require.Equal(t, tst.expected, result)
	}
}
