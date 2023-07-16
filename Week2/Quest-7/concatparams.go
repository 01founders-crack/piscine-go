package piscine

func ConcatParams(args []string) string {
	var myString string
	for i := 0; i < len(args); i++ {
		myString = myString + args[i]
		if i != len(args)-1 {
			myString = myString + "\n"
		}
	}
	return myString
}
