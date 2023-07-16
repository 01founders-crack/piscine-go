package piscine

func IterativePower(nb int, power int) int {
	if power == 0 {
		return 1
	}
	if power > 0 {
		result := 1
		for i := 0; i < power; i++ {
			result = result * nb
		}
		return result
	}
	return 0
}
