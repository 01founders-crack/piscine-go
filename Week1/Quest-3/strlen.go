package StrLen

func StrLen(s string) int {
	mArr := []rune(s)
	counter := 0
	for a := range mArr {
		counter++
		a++
	}
	return counter
}
