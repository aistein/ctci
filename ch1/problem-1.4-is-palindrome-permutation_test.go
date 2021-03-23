package ch1

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsPalindromePermutation(t *testing.T) {
	tests := []struct {
		name, s        string
		expected, skip bool
	}{
		{name: "Empty String", s: "", expected: true},
		{name: "Singleton", s: "a", expected: true},
		{name: "Double Is Perm", s: "aa", expected: true},
		{name: "Double Not Perm", s: "ab", expected: false},
		{name: "Short Mixed Case with Whitespace Is Perm", s: "T a t A", expected: true},
		{name: "Short Mixed Case with Whitespace Not Perm", s: "T a t C", expected: false},
		{name: "Mixed Case with Whitespace Is Perm", s: "Tact Coa", expected: true},
		// This test case must be skipped because of the overwhelming factorial runtime of the Brute Force solution
		{name: "Mixed Case with Whitespace Not Perm", s: "Tact Coa PQ RS", expected: false, skip: true},
	}
	implementations := []struct {
		name string
		fn   isPalindromePermutation
		skip bool
	}{
		{name: "Brute Force", fn: isPalindromePermutationBruteForce, skip: true},
		{name: "Hash Map", fn: isPalindromePermutationHashMap},
	}
	for _, implementation := range implementations {
		for _, tc := range tests {
			if !(tc.skip && implementation.skip) {
				t.Run(implementation.name+"/"+tc.name, func(t *testing.T) {
					require.Equal(t, tc.expected, implementation.fn(tc.s))
				})
			}
		}
	}
}

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name, s  string
		expected bool
	}{
		{name: "Empty String", s: "", expected: true},
		{name: "Singleton", s: "a", expected: true},
		{name: "Double Is Palindrome", s: "aa", expected: true},
		{name: "Double Not Palindrome", s: "ab", expected: false},
		{name: "Short Is Palindrome", s: "abcdefedcba", expected: true},
		{name: "Short Not Palindrome", s: "abcdefedbca", expected: false},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.expected, isPalindrome(tc.s))
		})
	}
}

func TestGeneratePermutations(t *testing.T) {
	tests := []struct {
		name, s  string
		expected map[string]bool
	}{
		{name: "Empty String", s: "", expected: map[string]bool{}},
		{name: "Singleton", s: "a", expected: map[string]bool{"a": true}},
		{name: "Double Same", s: "aa", expected: map[string]bool{"aa": true}},
		{name: "Double Unique", s: "ab", expected: map[string]bool{"ba": true, "ab": true}},
		{name: "Triple Unique", s: "abc", expected: map[string]bool{
			"abc": true,
			"acb": true,
			"bac": true,
			"bca": true,
			"cab": true,
			"cba": true,
		}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			permutations := generatePermutations(tc.s)
			require.Equal(t, len(tc.expected), len(permutations),
				"generated perms [%+v] do not match expected perms [%+v]", permutations, tc.expected)
			for perm := range permutations {
				if _, exists := tc.expected[perm]; !exists {
					t.Fatalf("generated permutation [%s] was not expected", perm)
				}
			}
		})
	}
}
