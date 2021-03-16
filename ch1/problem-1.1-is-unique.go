package ch1

import (
	"math/rand"
	"time"
)

type MutableString []rune
type IsUnique func(s *MutableString) bool

func isuniqueBruteforce(s *MutableString) bool {
	for i := 0; i < len(*s); i++ {
		for j := i+1; j < len(*s); j++ {
			if (*s)[i] == (*s)[j] {
				return false
			}
		}
	}
	return true
}

func isuniqueHashmap(s *MutableString) bool {
	seen := make(map[rune]bool)
	for _, runeVal := range *s {
		if _, exists := seen[runeVal]; exists {
			return false
		}
		seen[runeVal] = true
	}
	return true
}

func isuniqueQuicksort(s *MutableString) bool {
	if len(*s) < 2 {
		return true
	}

	quicksortMutableString(0, len(*s), s)

	prevRune := (*s)[0]
	for i := 1; i < len(*s); i++ {
		if (*s)[i] == prevRune {
			return false
		}
		prevRune = (*s)[i]
	}
	return true
}

func quicksortMutableString(l, h int, s *MutableString) {
	if l < h {
		j := pivotMutableString(l, h, s, func(l, h int) int {
			return l +rand.Intn(h - l)
		})
		quicksortMutableString(l, j, s)
		quicksortMutableString(j+1, h, s)
	}
}

// pivotIndexFunction must return an index in the range [l, h)
type pivotIndexFunction func(l, h int) int

// pivotMutableString selects a pivot from s using the supplied pivotIndexFunction; then it reorganizes s in place so
// that everything left of the pivot is less than the pivot, and everything after the pivot is greater than or equal to
// the pivot.  If s contains only unique values, the index where the pivot lands after reorganization is returned.  If s
// contains duplicate pivot values, the index where the rightmost duplicate of the pivot lands is returned.
func pivotMutableString(l, h int, s *MutableString, getIndex pivotIndexFunction) int {
	if h - l == 1 {
		return l
	}

	if h - l == 2 {
		if (*s)[l] > (*s)[l+1] {
			(*s)[l], (*s)[l+1] = (*s)[l+1], (*s)[l]
		}
		return l+1
	}

	pivotIdx := getIndex(l, h)
	pivot := (*s)[pivotIdx]

	lt, current, gt := l, l, h-1
	for current <= gt {
		if (*s)[current] < pivot {
			(*s)[current], (*s)[lt] = (*s)[lt], (*s)[current]
			current, lt = current+1, lt+1
			continue
		}
		if (*s)[current] == pivot {
			current++
			continue
		}
		if  (*s)[current] > pivot {
			(*s)[current], (*s)[gt] = (*s)[gt], (*s)[current]
			gt--
			continue
		}
	}
	return current - 1
}

func init() {
	// Quicksort will randomly select a pivot, so we want to intelligently seed the random number generator
	rand.Seed(time.Now().UnixNano())
}
