package ch1

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
	implementations := []struct {
		name string
		fn IsUnique
	}{
		{name: "Brute Force", fn: isuniqueBruteforce},
		{name: "Hash Map", fn: isuniqueHashmap},
		{name: "Quicksort", fn: isuniqueQuicksort},
	}
	for _, implementation := range implementations {
		for _, tc := range tests {
			t.Run(implementation.name + "/" + tc.name, func(t *testing.T){
				str := MutableString(tc.input)
				require.Equal(t, tc.expected, implementation.fn(&str))
			})
		}
	}
}

func TestPivotMutableString(t *testing.T) {
	tests := []struct {
		name string
		input string
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
		name string
		input string
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
