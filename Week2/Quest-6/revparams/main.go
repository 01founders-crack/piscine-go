package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args
	for i := len(arguments) - 1; i >= 1; i-- {
		printStringEndLine(arguments[i])
	}
}

func printStringEndLine(s string) {
	for _, v := range s {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}
