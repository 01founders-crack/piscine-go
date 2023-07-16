package piscine

func IsPrintable(str string) bool {
	runes := []rune(str)
	for i := range runes {
		if CheckPrint(runes[i]) == true {
			return false
		}
	}
	return true
}

func CheckPrint(x rune) bool {
	if x >= 0 && x <= 31 {
		return true
	}
	return false
}
