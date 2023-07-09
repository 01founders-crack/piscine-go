package StrRev

func StrRev(s string) string {
	mArr := []rune(s)
	var mArr2 string
	for i := len(s) - 1; i >= 0; i-- {
		mArr2 = mArr2 + string(mArr[i])
	}
	return mArr2
}
