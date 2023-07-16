package piscine

func Index(s string, toFind string) int {
	slength := len(s)
	toFindLen := len(toFind)

	t := 0
	for i := 0; i < slength; i++ {
		j := 0
		t = i
		for j < toFindLen {
			if t < slength && s[t] == toFind[j] {
				j++
				t++
			} else {
				break
			}
		}
		if j == toFindLen {
			return i
		}
	}
	return -1
}
