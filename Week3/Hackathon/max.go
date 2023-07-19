package piscine

func Max(a []int) int {
	highestvalue := 0
	for i := 0; i < len(a); i++ {
		if a[i] > highestvalue {
			highestvalue = a[i]
		}
	}
	return highestvalue
}
