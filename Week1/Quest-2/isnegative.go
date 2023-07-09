package IsNegative

import "github.com/01-edu/z01"

func IsNegative(nb int) {
	if nb >= 0 {
		z01.PrintRune(70)
		z01.PrintRune('\n')
	} else {
		z01.PrintRune(84)
		z01.PrintRune('\n')
	}
}
