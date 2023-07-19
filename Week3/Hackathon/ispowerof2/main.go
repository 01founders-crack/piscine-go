package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 2 {
		myInt, err := strconv.Atoi(os.Args[1])
		if err == nil {
			if powerOf2(myInt) {
				printString("true\n")
			} else {
				printString("false\n")
			}
		}

	}
}

func powerOf2(n int) bool {
	if n == 2 || n == 1 {
		return true
	} else {
		num := 2
		for i := 2; i < 4294967295; i += 2 {
			num = num * 2
			if num < n {
			} else if num == n {
				return true
			} else if num > n {
				return false
			}
		}
	}
	return false
}

func printString(s string) {
	for _, v := range s {
		z01.PrintRune(v)
	}
}
