package piscine

func SortWordArr(a []string) {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if compare2Strings(a[i], a[j]) == 1 {
				// if a[i] > a[j] {
				temp := a[i]
				a[i] = a[j]
				a[j] = temp
			}
		}
	}
}

func compare2Strings(s1, s2 string) int {
	s1Arr := StrToRuneArr(s1)
	s2Arr := StrToRuneArr(s2)
	minlen := 0
	if len(s1Arr) < len(s2Arr) {
		minlen = len(s1Arr)
	} else {
		minlen = len(s2Arr)
	}
	for i := 0; i < minlen; i++ {
		if s1Arr[i] > s2Arr[i] {
			return 1
		}
		if s1Arr[i] < s2Arr[i] {
			return -1
		}
	}
	if len(s1Arr) > len(s2Arr) {
		return 1
	} else if len(s1Arr) < len(s2Arr) {
		return -1
	} else {
		return 0
	}
}

func StrToRuneArr(s string) []rune {
	var mRune []rune
	if len(s) > 0 {
		for _, v := range s {
			mRune = append(mRune, v)
		}
	}
	return mRune
}
