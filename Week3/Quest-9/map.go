package piscine

func Map(f func(int) bool, a []int) []bool {
	var boolArray []bool
	for _, v := range a {
		boolArray = append(boolArray, f(v))
	}
	return boolArray
}
