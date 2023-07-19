package piscine

func Atoi(s string) int {
	myInt, err := mAtoi(s)
	if err == nil {
		return myInt
	} else {
		return 0
	}
}

func mAtoi(s string) (int, error) {
	for _, char := range s {
		if (char < 48 || char > 57) && char != 43 && char != 45 {
			return 0, nil
		}
	}
	mArr := []rune(s)
	isIt := false
	isItMinus := false

	for i, v := range mArr {
		if v == 43 || v == 45 {
			if i != 0 {
				return 0, nil
			} else {
				if v == 45 {
					isItMinus = true
				}
				isIt = true
			}
		}
	}
	if isIt {
		s = s[1:]
	}

	n := 0
	for _, ch := range s {
		ch -= '0'
		n = n*10 + int(ch)
	}
	if isItMinus {
		n = -n
	}
	return n, nil
	// Slow path for invalid, big, or underscored integers.
}
