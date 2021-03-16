package ch1

type CheckPermutation func(s1, s2 *MutableString) bool

func checkPermutationBruteForce(s1, s2 *MutableString) bool {
	if len(*s1) != len(*s2) {
		return false
	}
	for i := 0; i < len(*s1); i++ {
		cnt := 1
		for j := i+1; j < len(*s1); j++ {
			if (*s1)[j] == (*s1)[i] {
				cnt++
			}
		}
		for k := 0; k < len(*s2); k++ {
			if (*s2)[k] == (*s1)[i] {
				cnt--
			}
		}
		if cnt != 0 {
			return false
		}
	}
	return true
}

func checkPermutationHashMap(s1, s2 *MutableString) bool {
	if len(*s1) != len(*s2) {
		return false
	}
	counts := make(map[rune]int)
	for _, r := range *s1 {
		counts[r]++
	}
	for _, r := range *s2 {
		counts[r]--
		if counts[r] == 0 {
			delete(counts, r)
		}
	}
	return len(counts) == 0
}

func checkPermutationQuicksort(s1, s2 *MutableString) bool {
	if len(*s1) != len(*s2) {
		return false
	}
	quicksortMutableString(0, len(*s1), s1)
	quicksortMutableString(0, len(*s2), s2)
	for i := 0; i < len(*s1); i++ {
		if (*s1)[i] != (*s2)[i] {
			return false
		}
	}
	return true
}
