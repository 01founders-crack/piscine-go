package piscine

func ForEach(f func(int), arr []int) {
	for _, v := range arr {
		f(v)
	}
}
