package piscine

func CollatzCountdown(start int) int {
	counter := 0
	if start > 0 {
		for start > 1 {
			if start%2 == 0 {
				counter++
				start = start / 2
			} else {
				counter++
				start = (start * 3) + 1
			}
		}
	} else {
		return -1
	}
	return counter
}
