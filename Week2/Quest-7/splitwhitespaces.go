package piscine

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
