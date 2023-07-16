package piscine

func SortWordArr(a []string) {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if StrToRune(a[i]) > StrToRune(a[j])  {
				temp := a[i]
				a[i] = a[j]
				a[j] = temp
			}
		}
	}
}

func StrToRune(s string) rune {
    var mRune rune
    if len(s) > 0 {
        mRune = rune(s[0])
    }
    return mRune
}

