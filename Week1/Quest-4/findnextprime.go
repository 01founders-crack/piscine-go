package piscine

func FindNextPrime(nb int) int {
	for i := nb; i < 2147483647; i++ {
		if IsPrime(i) {
			return i
		}
	}
	return 0
}

func IsPrime(nb int) bool {
	if nb > 1 {
		for i := 2; i < int(nb/2)+1; i++ {
			if (nb % i) == 0 {
				return false
			}
		}
		return true
	} else {
		return false
	}
}
