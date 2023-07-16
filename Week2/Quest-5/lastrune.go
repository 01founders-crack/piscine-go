package piscine

func LastRune(s string) rune {
	a := rune(' ')
	for _, v := range s {
		a = rune(v)
	}
	return a
}
