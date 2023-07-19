package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	descending := true
	ascending := true

	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if f(a[i], a[j]) > 0 {
				ascending = false
			}
		}
	}
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if f(a[i], a[j]) <= 0 {
				descending = false
			}
		}
	}
	if ascending || descending {
		return true
	} else {
		return false
	}
}
