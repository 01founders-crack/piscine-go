package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrBase(number int, base string) {
	result := ""
	length := 0
	for _, char := range base {
		if char == char {
			length++
		}
	}
	maxPower := length
	if number < 0 {
		result = "-"
		maxPower *= -1
	}
	if length > 1 {
		for number/maxPower >= length {
			maxPower *= length
		}
		for maxPower != 0 {
			result = result + string(base[number/maxPower])
			number = number - number/maxPower*maxPower
			maxPower /= length
		}
		seen := map[rune]bool{}
		for _, char := range base {
			if char == '+' || char == '-' {
				result = "NV"
				break
			}
			if seen[char] == false {
				seen[char] = true
			} else {
				result = "NV"
				break
			}
		}
	} else {
		result = "NV"
	}
	for _, char := range result {
		z01.PrintRune(char)
	}
}
