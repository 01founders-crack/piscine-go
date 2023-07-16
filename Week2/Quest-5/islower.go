package piscine

func IsLower(s string) bool {
	for _, v := range s {
		if v >= 97 && v <= 122 {
		} else {
			return false
		}
	}
	return true
}
