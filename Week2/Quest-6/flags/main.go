package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) > 1 {
		args := os.Args[1:]

		IsOrder, IsHelp, IsInsert := checkArgs(args)
		// We checked the args and we know which ones we have.

		if IsHelp {
			printHelp()
		} else {
			mstr := returnPureString(args)
			if IsInsert {
				// z01.PrintRune('i')
				mstr = IsInsertAddToStr(args, mstr)
				// printString(mstr)
			}
			if IsOrder {
				printString(sortStringArrReturnStr(stringToStrArr(mstr)))
				printEndLine()
			} else {
				printString(mstr)
				printEndLine()
			}
		}
	} else {
		printHelp()
	}
}

func IsInsertAddToStr(args []string, mStr string) string {
	for _, v := range args {
		if len(v) > 9 {
			if v[0:9] == "--insert=" {
				mStr += v[9:]
			}
		}
		if len(v) > 3 {
			if v[0:3] == "-i=" {
				mStr += v[3:]
			}
		}

	}
	return mStr
}

func stringToStrArr(s string) []string {
	var myStringArr []string
	for _, v := range s {
		myStringArr = append(myStringArr, string(v))
	}
	return myStringArr
}

func returnPureString(args []string) string {
	var mstr string
	for _, v := range args {
		if len(v) > 0 {
			if v[0] != '-' {
				mstr = mstr + v
			}
		}
	}
	return mstr
}

func checkArgs(args []string) (bool, bool, bool) {
	IsOrder := false
	IsHelp := false
	IsInsert := false
	for _, v := range args {
		if v == "--order" || v == "-o" {
			IsOrder = true
		}
		if v == "--help" || v == "-h" {
			IsHelp = true
		}
		if v > "--insert=" || v > "-i=" {
			// Change the if because its different
			IsInsert = true
		}
	}
	return IsOrder, IsHelp, IsInsert
}

func printHelp() {
	printString("--insert\n")
	printString("  -i\n")
	printString("	 This flag inserts the string into the string passed as argument.\n")
	printString("--order\n")
	printString("  -o\n")
	printString("	 This flag will behave like a boolean, if it is called it will order the argument.\n")
}

func sortStringArrReturnStr(args []string) string {
	for i := 0; i < len(args); i++ {
		for j := i + 1; j < len(args); j++ {
			if compare(args[j], args[i]) {
				args[i], args[j] = args[j], args[i]
			}
		}
	}
	var mString string
	for _, v := range args {
		mString += v
	}
	return mString
}

func compare(a, b string) bool {
	for i := 0; i < len(a) && i < len(b); i++ {
		if a[i] != b[i] {
			return a[i] < b[i]
		}
	}
	return len(a) < len(b)
}

func printString(s string) {
	for _, v := range s {
		z01.PrintRune(v)
	}
}

func printEndLine() {
	z01.PrintRune('\n')
}
