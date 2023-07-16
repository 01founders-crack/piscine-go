package main

import "os"

func main() {
	if len(os.Args) == 4 {
		mArgs := os.Args[1:]
		if isInt64(mArgs[0]) && isInt64(mArgs[2]) {
			if mArgs[1] == "+" {

			}
			if mArgs[1] == "-" {

			}
			if mArgs[1] == "/" {
				if mArgs[2] == "0" {

				}
			}
			if mArgs[1] == "*" {

			}
			if mArgs[1] == "%" {
				if mArgs[2] == "0" {

				}
			}
		}
	}
}

func isInt64(s string) bool {
	return true
}

func StrToRune(s string) rune {
	var mRune rune
	if len(s) > 0 {
		mRune = rune(s[0])
	}
	return mRune
}
