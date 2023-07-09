package piscine

func QuadB(a, b int) {
	for x := 0; x < b; x++ {
		for y := 0; y < a; y++ {
			if x == 0 {
				IfFirstLine(x, y, a, b)
			}
			if x > 0 && x < b-1 {
				IfMidLine(x, y, a, b)
			}
			if x == b-1 && b != 1 {
				IfLastLine(x, y, a, b)
			}
		}
		NoNegative(a, b)
	}
}

func IfFirstLine(i, j, a, b int) {
	if j == 0 && i == 0 {
		print("/")
	}
	if j > 0 && j < a-1 {
		print("*")
	}
	if j == a-1 && i == 0 {
		if a > 1 {
			print("\\")
		} else if b != 1 && a != 1 {
			print("\\")
		}
	}
}

func IfMidLine(i, j, a, b int) {
	if i > 0 && i < b-1 {
		if j == 0 || j == a-1 {
			print("*")
		}
		if j > 0 && j < a-1 {
			print(" ")
		}
	}

}

func IfLastLine(i, j, a, b int) {
	if j == 0 && i == b-1 {
		if a != 1 {
			print("\\")
		} else if b > 1 {
			print("\\")
		}

	}
	if j > 0 && j < a-1 {
		print("*")
	}
	if j == a-1 && i == b-1 && a != 1 {
		print("/")
	}
}

func NoNegative(a, b int) {
	if a >= 0 && b >= 0 {
		print("\n")
	}
}
