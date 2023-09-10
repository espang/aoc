package aoc

import (
	"errors"
	"strconv"
	"unicode"
	"unicode/utf8"
)

func ParseLetterNumber(v string) (rune, int, error) {
	if len(v) == 0 {
		return 0, 0, errors.New("expect non empty string")
	}

	r, size := utf8.DecodeRuneInString(v)
	if size != 1 || !unicode.IsLetter(r) {
		return 0, 0, errors.New("expect first rune to be a letter")
	}

	number, err := strconv.ParseInt(v[1:], 10, 64)
	if err != nil {
		return 0, 0, err
	}
	return r, int(number), nil
}
