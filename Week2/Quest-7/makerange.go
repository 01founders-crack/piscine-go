package piscine

func MakeRange(min, max int) []int {
	if min < max {
		mArr := make([]int, max-min)
		counter := 0
		for i := min; i < max; i++ {
			mArr[counter] = i
			counter++
		}
		return mArr
	} else {
		var Arr []int
		return Arr
	}
}
