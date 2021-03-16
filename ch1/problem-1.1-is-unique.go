package ch1

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

