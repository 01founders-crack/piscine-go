package piscine

func IsNumeric(s string) bool {
	for _, v := range s {
		if v >= 48 && v <= 57 {
		} else {
			return false
		}
	}
	return true
}
