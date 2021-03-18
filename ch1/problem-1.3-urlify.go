package ch1

type urlify func(ms *MutableString, trueLength int)

func urlifyBruteForce(ms *MutableString, trueLength int) {
	var i int
	for i < trueLength {
		if (*ms)[i] == ' ' {
			for j := trueLength - 1; j > i+2; j-- {
				(*ms)[j] = (*ms)[j-2]
			}
			(*ms)[i], (*ms)[i+1], (*ms)[i+2] = '%', '2', '0'
			i += 3
		} else {
			i++
		}
	}
}

func urlifyBackwards(ms *MutableString, trueLength int) {
	if trueLength > 0 {
		curr, target := trueLength-1, trueLength-1
		for (*ms)[curr] == -1 {
			curr--
		}
		for curr >= 0 {
			if (*ms)[curr] == ' ' {
				(*ms)[target-2], (*ms)[target-1], (*ms)[target] = '%', '2', '0'
				target -= 3
			} else {
				(*ms)[target] = (*ms)[curr]
				target--
			}
			curr--
		}
	}
}
