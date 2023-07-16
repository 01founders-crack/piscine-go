package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := (os.Args)[0]
	for _, v := range arguments[2:] {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}

func reversePrintStringz01(s string) {
	b := []rune(s)
	for i := len(b) - 1; i >= 0; i-- {
		z01.PrintRune(b[i])
	}
}

func splitString(s string) string {
	var a string
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != '/' {
			a = a + string(s[i])
		} else {
			return a
		}
	}
	return a
}
