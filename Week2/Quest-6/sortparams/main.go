package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	sortStringArr(args)
	printStringArr(args)
}

func sortStringArr(args []string) {
	for i := 0; i < len(args); i++ {
		for j := i + 1; j < len(args); j++ {
			if compare(args[j], args[i]) {
				args[i], args[j] = args[j], args[i]
			}
		}
	}
}

func compare(a, b string) bool {
	for i := 0; i < len(a) && i < len(b); i++ {
		if a[i] != b[i] {
			return a[i] < b[i]
		}
	}
	return len(a) < len(b)
}

func printStringWz01(s string) {
	for _, v := range s {
		z01.PrintRune(v)
	}
}

func printStringArr(args []string) {
	for _, arg := range args {
		printStringWz01(arg)
		z01.PrintRune('\n')
	}
}
