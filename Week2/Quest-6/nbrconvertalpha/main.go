package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) > 1 {
		args := os.Args[1:]
		if args[0] == "--upper" {
			args = args[1:]
			argsIntArr := StrArrToIntArr(args)
			printIntArrToLetter(argsIntArr, 64)

		} else {
			argsIntArr := StrArrToIntArr(args)
			printIntArrToLetter(argsIntArr, 96)
		}
		z01.PrintRune('\n')
	}
}

func printIntArrToLetter(x []int, n int) {
	for _, v := range x {
		if v > 0 && v < 27 {
			z01.PrintRune(rune(v + n))
		} else {
			z01.PrintRune(' ')
		}
	}
}

func StrArrToIntArr(s []string) []int {
	var mIntArr []int
	for _, bb := range s {
		numv := 0
		for _, v := range bb {
			numv = numv*10 + int(v-'0')
		}
		mIntArr = append(mIntArr, numv)
	}
	return mIntArr
}
