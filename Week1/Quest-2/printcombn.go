package PrintNbr

import (
	"github.com/01-edu/z01"
)

func PrintCombN(n int) {
	for i := 48; i < 58; i++ {
		for j := 48; j < 58; j++ {
			for x := 48; x < 58; x++ {
				for y := 48; y < 58; y++ {
					if i == x {
						if j < y {
							PrintNumbers(i, j, x, y)
							IfPrintNumbers(i, j, x, y)
						}
					} else if i < x {
						PrintNumbers(i, j, x, y)
						PrintCommaSpace()
					}
				}
			}
		}
	}
}

func PrintNumbers(i int, j int, x int, y int) {
	z01.PrintRune(rune(i))
	z01.PrintRune(rune(j))
	z01.PrintRune(' ')
	z01.PrintRune(rune(x))
	z01.PrintRune(rune(y))
}

func IfPrintNumbers(i int, j int, x int, y int) {
	if i == 57 && j == 56 {
		z01.PrintRune('\n')
	} else {
		PrintCommaSpace()
	}
}

func PrintCommaSpace() {
	z01.PrintRune(',')
	z01.PrintRune(' ')
}
