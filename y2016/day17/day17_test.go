package day17

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHash(t *testing.T) {
	label := func(in, out string) string { return fmt.Sprintf("md5(%s)=%s", in, out) }
	hasher := NewHasher()
	for _, tc := range []struct {
		in  string
		out string
	}{
		{"hijkl", "ced9"},
		{"hijklD", "f2bc"},
		{"hijklDU", "528e"},
	} {
		t.Run(label(tc.in, tc.out), func(t *testing.T) {
			assert.Equal(t, []byte(tc.out), hasher.HashOf([]byte(tc.in)))
		})
	}
}

// TestDay17 to make sure that the code will not be broken by changes in any shared code.
func TestDay17(t *testing.T) {
	t.Run("Part1 returns the same result", func(t *testing.T) {
		result, err := Part1()
		require.NoError(t, err)
		assert.Equal(t, "RDURRDDLRD", result)
	})

	t.Run("Part2 returns the same result", func(t *testing.T) {
		result, err := Part2()
		require.NoError(t, err)
		assert.Equal(t, "526", result)
	})
}
