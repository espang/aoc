package aoc_test

import (
	"testing"

	"github.com/espang/aoc/aoc"
)

func equal(t *testing.T, expected, actual any) {
	if expected == actual {
		return
	}
	t.Errorf("expected %v and got %v", expected, actual)
}

func TestParseLetterNumber(t *testing.T) {
	t.Run("good numbers", func(t *testing.T) {
		letter, number, err := aoc.ParseLetterNumber("L90")
		if err != nil {
			t.Fatal(err)
		}
		equal(t, 'L', letter)
		equal(t, 90, number)
	})

	t.Run("empty string should fail", func(t *testing.T) {
		_, _, err := aoc.ParseLetterNumber("")
		if err == nil {
			t.Fatal("expeceted error but it was nil")
		}
	})
	t.Run("when the first rune is not a letter", func(t *testing.T) {
		_, _, err := aoc.ParseLetterNumber("-90")
		if err == nil {
			t.Fatal("expeceted error but it was nil")
		}
	})
	t.Run("when the rest of the string is not a number", func(t *testing.T) {
		_, _, err := aoc.ParseLetterNumber("L90E")
		if err == nil {
			t.Fatal("expeceted error but it was nil")
		}
	})
}
