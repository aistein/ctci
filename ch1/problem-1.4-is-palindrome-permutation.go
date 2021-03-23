package ch1

import (
	"unicode"
)

type isPalindromePermutation func(s string) bool

func isPalindromePermutationBruteForce(s string) bool {
	runes := []rune(s)
	if len(runes) == 0 {
		return true
	}
	var filteredRunes []rune
	for _, r := range runes {
		if unicode.IsLetter(r) {
			filteredRunes = append(filteredRunes, unicode.ToLower(r))
		}
	}
	for perm := range generatePermutations(string(filteredRunes)) {
		if isPalindrome(perm) {
			return true
		}
	}
	return false
}

func isPalindromePermutationHashMap(s string) bool {
	unpaired := make(map[rune]bool)
	for _, r := range s {
		if unicode.IsLetter(r) {
			r = unicode.ToLower(r)
		} else {
			continue
		}
		if _, exists := unpaired[r]; exists {
			delete(unpaired, r)
		} else {
			unpaired[r] = true
		}
	}
	return len(unpaired) <= 1
}

func generatePermutations(s string) map[string]bool {
	permutations := make(map[string]bool)
	runes := []rune(s)

	// Base Case
	if len(runes) == 1 {
		return map[string]bool{string(runes): true}
	}

	// Recursive Case
	for i := 0; i < len(runes); i++ {
		for partial := range generatePermutations(string(runes[:i]) + string(runes[i+1:])) {
			permutations[string(runes[i])+partial] = true
		}
	}

	return permutations
}

func isPalindrome(s string) bool {
	runes := []rune(s)
	for i := 0; i < len(runes)/2; i++ {
		if runes[i] != runes[len(runes)-i-1] {
			return false
		}
	}
	return true
}
