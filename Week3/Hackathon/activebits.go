package piscine

func ActiveBits(n int) int {
	counter := 0
	calculate := lpowerOf2(n)
	numb := n
	for numb >= 0 {
		if numb >= calculate {
			numb -= calculate
			counter++
		} else if calculate > 1 {
			calculate = calculate / 2
		} else {
			numb -= 1
		}
	}
	return counter
}

func lpowerOf2(n int) int {
	num := 1
	if n == 1 || n == 2 {
		return n
	} else {
		for num < n {
			num = num * 2
		}
	}
	return num
}
