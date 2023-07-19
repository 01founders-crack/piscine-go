package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 4 {
		mArgs := os.Args[1:]
		mArgs0, err0 := strconv.Atoi(mArgs[0])
		mArgs2, err2 := strconv.Atoi(mArgs[2])
		if err0 == nil && err2 == nil {
			if mArgs[1] == "+" {
				sum02, errAdd64 := Add64(mArgs0, mArgs2)
				if errAdd64 {
					printInt(sum02)
				}
			}
			if mArgs[1] == "-" {
				diff02, errSub64 := Sub64(mArgs0, mArgs2)
				if errSub64 {
					printInt(diff02)
				}
			}
			if mArgs[1] == "/" {
				if mArgs2 == 0 {
					printString("No division by 0\n")
				} else {
					div02, errDiv64 := Div64(mArgs0, mArgs2)
					if errDiv64 {
						printInt(div02)
					}
				}
			}
			if mArgs[1] == "*" {
				mul02, errMul64 := Mul64(mArgs0, mArgs2)
				if errMul64 {
					printInt(mul02)
				}
			}
			if mArgs[1] == "%" {
				if mArgs2 == 0 {
					printString("No modulo by 0\n")
				} else {
					mod02 := mArgs0 % mArgs2
					printInt(mod02)
				}
			}
		}
	}
}

func printInt(mint int) {
	isMinus, abc := extractDigits(mint)
	if isMinus {
		z01.PrintRune('-')
	}
	for i := len(abc) - 1; i >= 0; i-- {
		z01.PrintRune(abc[i] + 48)
	}
	z01.PrintRune('\n')
}

func printString(s string) {
	for _, v := range s {
		z01.PrintRune(v)
	}
}

func StrToRune(s string) rune {
	var mRune rune
	if len(s) > 0 {
		mRune = rune(s[0])
	}
	return mRune
}

func extractDigits(n int) (bool, []rune) {
	if n == 0 {
		return true, []rune{0}
	}
	var digits []rune
	isMinus := false
	if n < 0 {
		n = -n
		isMinus = true
	}
	for n > 0 {
		digit := n % 10
		digits = append(digits, rune(digit))
		n /= 10
	}
	return isMinus, digits
}

func Add64(a, b int) (int, bool) {
	c := a + b
	if (c > a) == (b > 0) {
		return c, true
	}
	return c, false
}

func Sub64(a, b int) (int, bool) {
	c := a - b
	if (c < a) == (b > 0) {
		return c, true
	}
	return c, false
}

// Mul64 performs * operation on two int64 operands
// returning a result and status
func Mul64(a, b int) (int, bool) {
	if a == 0 || b == 0 {
		return 0, true
	}
	c := a * b
	if (c < 0) == ((a < 0) != (b < 0)) {
		if c/b == a {
			return c, true
		}
	}
	return c, false
}

// Quotient64 performs + operation on two int64 operands
// returning a quotient, a remainder and status
func Quotient64(a, b int) (int, int, bool) {
	if b == 0 {
		return 0, 0, false
	}
	c := a / b
	status := (c < 0) == ((a < 0) != (b < 0))
	return c, a % b, status
}

// Div64 performs / operation on two int operands
// returning a result and status
func Div64(a, b int) (int, bool) {
	q, _, ok := Quotient64(a, b)
	return q, ok
}
