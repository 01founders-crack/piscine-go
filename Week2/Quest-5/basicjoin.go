package piscine

func BasicJoin(elems []string) string {
	myString := ""
	for _, v := range elems {
		myString = myString + v
	}
	return myString
}
