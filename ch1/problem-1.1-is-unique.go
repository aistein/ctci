package ch1

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
