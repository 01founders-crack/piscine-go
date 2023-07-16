package piscine

func IsUpper(s string) bool {
	for _, v := range s {
		if v >= 65 && v <= 90 {
		} else {
			return false
		}
	}
	return true
}
