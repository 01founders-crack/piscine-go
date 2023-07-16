package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	digits := extractDigits(n)
	sortDigits(digits)
	printDigits(digits)
}

func extractDigits(n int) []int {
	if n == 0 {
		return []int{0}
	}

	var digits []int

	for n > 0 {
		digit := n % 10
		digits = append(digits, digit)
		n /= 10
	}

	return digits
}

func sortDigits(digits []int) {
	length := len(digits)

	for i := 0; i < length-1; i++ {
		for j := 0; j < length-i-1; j++ {
			if digits[j] > digits[j+1] {
				digits[j], digits[j+1] = digits[j+1], digits[j]
			}
		}
	}
}

func printDigits(digits []int) {
	for _, digit := range digits {
		z01.PrintRune(rune(digit + 48))
	}
}
