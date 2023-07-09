package PrintCombN

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	num := make([]int, n)
	printCombinations(num, 0, 0)
}

func printCombinations(num []int, index int, start int) {
	if index == len(num) {
		printNumber(num)
		return
	}

	for i := start; i <= 9; i++ {
		num[index] = i
		printCombinations(num, index+1, i+1)
	}
}

func printNumber(num []int) {
	for i := 0; i < len(num); i++ {
		z01.PrintRune(rune('0' + num[i]))
		if i != len(num)-1 {
			printCommaSpace()
		}
	}
	printNewLine()
}

func printCommaSpace() {
	z01.PrintRune(',')
	z01.PrintRune(' ')
}

func printNewLine() {
	z01.PrintRune('\n')
}
