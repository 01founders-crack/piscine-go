package BasicAtoi2

func BasicAtoi2(s string) int {
	myInt, err := Atoi(s)
	if err == nil {
		return myInt
	} else {
		return 0
	}
}

func Atoi(s string) (int, error) {
	for _, char := range s {
		if char < 48 || char > 57 {
			return 0, nil
		}
	}
	n := 0
	for _, ch := range []byte(s) {
		ch -= '0'
		n = n*10 + int(ch)
	}
	return n, nil
	// Slow path for invalid, big, or underscored integers.
}
