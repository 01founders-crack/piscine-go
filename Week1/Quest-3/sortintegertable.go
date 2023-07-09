package piscine

func SortIntegerTable(table []int) {
	for i := 0; i < len(table); i++ {
		for j := len(table) - 1; j > i; j-- {
			if table[j] < table[j-1] {
				table[j], table[j-1] = table[j-1], table[j]
			}
		}
	}
}
