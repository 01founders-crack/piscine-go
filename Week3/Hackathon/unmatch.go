package piscine

func Unmatch(a []int) int {
	for i := 0; i < len(a); i++ {
		counter := 0
		for j := 0; j < len(a); j++ {
			if i != j && a[i] == a[j] {
				counter++
			}
		}
		if counter == 0 || counter%2 == 0 {
			return a[i]
		}
	}
	return -1
}
