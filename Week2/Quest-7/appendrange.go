package piscine

func AppendRange(min, max int) []int {
	var mArr []int

	if min < max {
		for i := min; i < max; i++ {
			mArr = append(mArr, i)
		}
	}
	return mArr
}
