package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for _, v := range a {
		printString(v)
		z01.PrintRune('\n')
	}
}

func printString(s string) {
	for _, v := range s {
		z01.PrintRune(v)
	}
}

func SplitWhiteSpaces(str string) []string {
	splitStr := make([]string, 0)
	word := ""

	for _, char := range str {
		if char != ' ' {
			word += string(char)
		} else {
			if word != "" {
				splitStr = append(splitStr, word)
			}
			word = ""
		}
	}

	// Include the last word after the last space
	if word != "" {
		splitStr = append(splitStr, word)
	}
	return splitStr
}
