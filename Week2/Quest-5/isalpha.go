package piscine

func IsAlpha(s string) bool {
	for _, v := range s {
		if v >= 97 && v <= 122 || v >= 65 && v <= 90 || v >= 48 && v <= 57 {
		} else {
			return false
		}
	}
	return true
}
