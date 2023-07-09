package main

import "github.com/01-edu/z01"

func main() {
	for i := 122; i > 96; i-- {
		z01.PrintRune(rune(i))
	}
	z01.PrintRune('\n')
}
