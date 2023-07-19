package piscine

func Rot14(s string) string {
	finalstring := ""
	runelist := []rune(s)
	for _, r := range runelist {
		finalstring += string(Next14(r))
	}
	return finalstring
}

func Next14(r rune) rune {
	counter := 14
	if r >= 65 && r <= 90 {
		for i := 0; i < counter; i++ {
			if r < 90 {
				r = r + 1
			} else {
				r = 65
			}
		}
	} else if r >= 97 && r <= 122 {
		for i := 0; i < counter; i++ {
			if r < 122 {
				r = r + 1
			} else {
				r = 97
			}
		}
	}
	return r
}
