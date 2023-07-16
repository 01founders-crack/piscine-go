package piscine

func AlphaCount(s string) int {
	counter := 0
	for _, v := range s {
		if rune(v) <= rune(90) && rune(v) >= rune(65) {
			counter++
		}
		if rune(v) <= rune(122) && rune(v) >= rune(97) {
			counter++
		}
	}
	return counter
}
