package main

import (
	"errors"
	"fmt"
	"os"
	"unicode"
)

func main() {
	s := os.Args
	if len(s) > 2 {
		_, _ = fmt.Fprintln(os.Stderr, "Too many args in command line")
		return
	}
	fmt.Println(Unpack("\\55qwe\\\\5"))
}

var errInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	sr := []rune(s)
	var s2 string
	var n int
	var backslash bool

	for i, item := range sr {
		if unicode.IsDigit(item) && i == 0 {
			return "", errInvalidString
		}

		if unicode.IsDigit(item) && unicode.IsDigit(sr[i-1]) && sr[i-2] != '\\' {
			return "", errInvalidString
		}
		if item == '\\' && !backslash {
			backslash = true
			continue
		}
		if backslash && unicode.IsLetter(item) {
			return "", errInvalidString
		}
		if backslash {
			s2 += string(item)
			backslash = false
			continue
		}
		if unicode.IsDigit(item) {
			n = int(item - '0')
			if n == 0 {
				s2 = s2[:len(s2)-1]
				continue
			}
			for j := 0; j < n-1; j++ {
				s2 += string(sr[i-1])
			}
			continue
		}
		s2 += string(item)
	}

	return s2, nil
}
