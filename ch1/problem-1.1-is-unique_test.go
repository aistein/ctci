package ch1

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsUnique(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{name: "Empty", input: "", expected: true},
		{name: "Singleton", input: "a", expected: true},
		{name: "Double", input: "aa", expected: false},
		{name: "Short Unique", input: "abcdefg", expected: true},
		{name: "Short Not Unique", input: "abcdafg", expected: false},
		{name: "Mixed Long Unique", input: "aBcdefghijkLmnopqRstuvWxyz@!+=#2^", expected: true},
		{name: "Mixed Long Not Unique", input: "saiouaFASKDFjasdfhkli$^$%E(+++", expected: false},
	}
	implementations := []struct {
		name string
		fn   isUnique
	}{
		{name: "Brute Force", fn: isuniqueBruteforce},
		{name: "Hash Map", fn: isuniqueHashmap},
		{name: "Quicksort", fn: isuniqueQuicksort},
	}
	for _, implementation := range implementations {
		for _, tc := range tests {
			t.Run(implementation.name+"/"+tc.name, func(t *testing.T) {
				str := MutableString(tc.input)
				require.Equal(t, tc.expected, implementation.fn(&str))
			})
		}
	}
}
