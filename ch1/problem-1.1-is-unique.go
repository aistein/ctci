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
