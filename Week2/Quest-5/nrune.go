package piscine

func NRune(s string, n int) rune {
	if n < 1 {
		return rune(0)
	} else {
		for index, v := range s {
			if index+1 == n {
				return rune(v)
			}
		}
	}
	return rune(0)
}
