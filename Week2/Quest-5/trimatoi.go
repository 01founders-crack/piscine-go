package piscine

func TrimAtoi(s string) int {
	var result int
	sign := 1
	foundNum := false

	for _, letter := range s {
		// Accumulate digits only
		if letter >= '0' && letter <= '9' {
			result = result*10 + int(letter-'0')
			foundNum = true
		} else if letter == '-' && !foundNum {
			// Handle negative sign only if no digit has been found yet
			sign = -1
		}
	}
	return result * sign
}
