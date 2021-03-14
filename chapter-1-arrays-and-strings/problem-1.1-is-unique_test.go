package chapter_1_arrays_and_strings

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIsUnique(t *testing.T) {
	tests := []struct {
		name string
		input string
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
	implementations := []IsUnique{
		isuniqueBruteforce,
	}
	for _, implementation := range implementations {
		for _, tc := range tests {
			t.Run(tc.name, func(t *testing.T){
				str := MutableString(tc.input)
				require.Equal(t, tc.expected, implementation(&str))
			})
		}
	}
}
