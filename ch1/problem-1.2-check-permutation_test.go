package ch1

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCheckPermutation(t *testing.T) {
	tests := []struct{
		name, s1, s2 string
		expected bool
	}{
		{name: "Empty Strings", s1: "", s2: "", expected: true},
		{name: "Singletons Match", s1: "a", s2: "a", expected: true},
		{name: "Singletons Mismatch", s1: "a", s2: "b", expected: false},
		{name: "Short Match", s1: "abc", s2: "bac", expected: true},
		{name: "Short Mismatch", s1: "abc", s2: "bad", expected: false},
		{name: "Short Mismatch Length", s1: "abc", s2: "abbc", expected: false},
		{name: "Long Match", s1: "abcdefghijklmnop123#$%", s2: "#$%abcdefg123hijklmnop", expected: true},
		{name: "Long Mismatch", s1: "abcdefghijklmnop123#$%", s2: "#$%abcdefg123hijklmnopq", expected: false},
	}
	implementations := []struct {
		name string
		fn CheckPermutation
	}{
		{name: "Brute Force", fn: checkPermutationBruteForce},
		{name: "Hash Map", fn: checkPermutationHashMap},
		{name: "Quicksort", fn: checkPermutationQuicksort},
	}
	for _, implementation := range implementations {
		for _, tc := range tests {
			t.Run(implementation.name + "/" + tc.name, func(t *testing.T) {
				s1, s2 := MutableString(tc.s1), MutableString(tc.s2)
				require.Equal(t, tc.expected, implementation.fn(&s1, &s2))
			})
		}
	}
}
