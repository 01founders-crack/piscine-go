package piscine

func Join(strs []string, sep string) string {
	myString := ""
	for index, v := range strs {
		myString = myString + v
		if index != len(strs)-1 {
			myString = myString + sep
		}
	}
	return myString
}
