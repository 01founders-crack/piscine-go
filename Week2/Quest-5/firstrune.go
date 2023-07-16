package piscine

func FirstRune(s string) rune {
	a := rune(' ')
	for _, v := range s {
		a = rune(v)
		break
	}
	return a
}
