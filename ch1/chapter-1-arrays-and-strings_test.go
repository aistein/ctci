package ch1

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPivotMutableString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{name: "Singleton", input: "a", expected: 0},
		{name: "Double Unique", input: "ab", expected: 1},
		{name: "Double Not Unique", input: "aa", expected: 1},
		{name: "Short Unique", input: "dbcaefg", expected: 3},
		{name: "Short Not Unique", input: "dbcadfg", expected: 4},
		{name: "Long Unique", input: "xyzwabqfglnmpr", expected: 11},
		{name: "Long Not Unique", input: "xyzwabxqfglnmpfr", expected: 13},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ms := MutableString(tc.input)
			getStablePivotIndex := func(l, h int) int {
				return l
			}
			require.Equal(t, tc.expected, pivotMutableString(0, len(ms), &ms, getStablePivotIndex))
		})
	}
}

func TestQuicksortMutableString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "Singleton", input: "a", expected: "a"},
		{name: "Double Unique", input: "ba", expected: "ab"},
		{name: "Double Not Unique", input: "aa", expected: "aa"},
		{name: "Short Unique", input: "dbcaefg", expected: "abcdefg"},
		{name: "Short Not Unique", input: "dbcadfg", expected: "abcddfg"},
		{name: "Long Unique", input: "xyzwabqfglnmpr", expected: "abfglmnpqrwxyz"},
		{name: "Long Not Unique", input: "xyzwabqfglnmpfr", expected: "abffglmnpqrwxyz"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ms := MutableString(tc.input)
			quicksortMutableString(0, len(ms), &ms)
			require.Equal(t, tc.expected, string(ms))
		})
	}
}
