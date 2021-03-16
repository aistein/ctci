package ch1

import (
	"math/rand"
	"time"

	"github.com/sirupsen/logrus"
)

// MutableString is a proxy for []rune.  Since the built-in "string" type is an immutable []byte in golang, pre-casting
// "string" into []rune abstracts away some of the golang-specific implementation details from string-mutating
// algorithms.  Using []rune also bakes in the assumption that "MutableStrings can have every possible character
// encoding, even mixing multiple encodings."
type MutableString []rune

func ChapterName() {
	logrus.Println("Arrays and Strings")
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
