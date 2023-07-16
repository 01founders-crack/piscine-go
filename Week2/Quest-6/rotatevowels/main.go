package main

import (
	"github.com/01-edu/z01"
)

func main() {
}

func printString(s string) {
	for _, v := range s {
		z01.PrintRune(v)
	}
}

func printEndLine() {
	z01.PrintRune('\n')
}
