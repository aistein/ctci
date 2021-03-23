package ch1

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestURLify(t *testing.T) {
	tests := []struct {
		name, s, expected string
		trueLength        int
	}{
		{name: "Empty String", s: "", expected: "", trueLength: 0},
		{name: "Single Whitespace", s: " ", expected: "%20", trueLength: 3},
		{name: "Single Non-Whitespace", s: "a", expected: "a", trueLength: 1},
		{name: "Multiple Whitespace", s: "   ", expected: "%20%20%20", trueLength: 9},
		{name: "Multiple Non-Whitespace", s: "abcd", expected: "abcd", trueLength: 4},
		{name: "Mixed Whitespace and Non-Whitespace", s: "Mr John Smith", expected: "Mr%20John%20Smith", trueLength: 17},
	}
	implementations := []struct {
		name string
		fn   urlify
	}{
		{name: "Brute Force", fn: urlifyBruteForce},
		{name: "Backwards", fn: urlifyBackwards},
	}
	for _, implementation := range implementations {
		for _, tc := range tests {
			t.Run(implementation.name+"/"+tc.name, func(t *testing.T) {
				ms := toMutableStringWithCapacity(tc.s, tc.trueLength)
				implementation.fn(&ms, tc.trueLength)
				require.Equal(t, tc.expected, string(ms))
			})
		}
	}
}

func toMutableStringWithCapacity(s string, c int) MutableString {
	ms := MutableString(s)
	for i := 0; i < c-len(s); i++ {
		ms = append(ms, -1)
	}
	return ms
}
